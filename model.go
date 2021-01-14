package main

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

type Model struct {
	hasher hash.Hash

	providedHash string
	actualHash   string
	filePath     string
	hashType     string

	ctx          context.Context
	cancelFunc   context.CancelFunc
	resultFunc   func(bool, error)
	progressFunc func(float32)
}

func NewModel() *Model {
	return &Model{}
}

func (model *Model) DetectProvidedHashType() string {
	switch len(model.providedHash) {
	case md5.Size * 2:
		model.hasher = md5.New()
		model.hashType = "MD5"
	case sha256.Size * 2:
		model.hasher = sha256.New()
		model.hashType = "SHA-256"
	case sha512.Size * 2:
		model.hasher = sha512.New()
		model.hashType = "SHA-512"
	default:
		model.hasher = nil
		model.hashType = ""
	}

	return model.hashType
}

func (model *Model) IsGoodToGo() bool {
	return model.hashType != "" && model.filePath != "" && model.providedHash != ""
}

func (model *Model) CreateContext() {
	model.ctx, model.cancelFunc = context.WithCancel(context.Background())
}

func (model *Model) StopHashing() {
	model.cancelFunc()
}

func (model *Model) StartHashing() {
	// Open file
	file, err := os.OpenFile(model.filePath, os.O_RDONLY, 0666)
	if err != nil {
		return
	}

	stat, err := file.Stat()
	if err != nil {
		return
	}

	// On exit
	defer func() {
		model.actualHash = hex.EncodeToString(model.hasher.Sum(nil))
		model.hasher.Reset()
		model.resultFunc(model.actualHash == model.providedHash, err)
		model.cancelFunc()
		file.Close()
	}()

	//bufferSize := int64(32 * 1024)
	fileSize := stat.Size()
	processedBytes := int64(0)
	bufferSize := fileSize / 100

	// Read file
	for {
		// Check if context were cancelled
		select {
		case <-model.ctx.Done():
			err = model.ctx.Err()
			return
		default:
		}

		// Define buffer
		buffer := make([]byte, bufferSize)

		// Read bytes
		readBytes, err := file.Read(buffer)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
		buffer = buffer[:readBytes]

		// Write bytes to hasher
		readBytes, err = model.hasher.Write(buffer)
		if err != nil {
			return
		}
		processedBytes += int64(readBytes)

		model.progressFunc(float32(processedBytes) / float32(fileSize))
	}
}
