package model

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

	totalSize   int64
	currentSize int64

	ctx        context.Context
	cancelFunc context.CancelFunc
	resultFunc func(bool, error)
}

func New() *Model {
	return &Model{}
}

func (model *Model) SetFile(f string) {
	model.filePath = f
}

func (model *Model) SetHash(h string) {
	model.providedHash = h
}

func (model *Model) SetResultFunc(f func(bool, error)) {
	model.resultFunc = f
}

func (model *Model) GetProgress() float64 {
	return float64(model.currentSize) / float64(model.totalSize)
}

func (model *Model) DetectType() string {
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

func (model *Model) IsReady() bool {
	return model.hashType != "" && model.filePath != "" && model.providedHash != ""
}

func (model *Model) PrepareHashing() {
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

	model.totalSize = stat.Size()
	model.currentSize = 0

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
		buffer := make([]byte, 32*1024)

		// Read bytes
		n, err := file.Read(buffer)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
		buffer = buffer[:n]

		// Write bytes to hasher
		n, err = model.hasher.Write(buffer)
		if err != nil {
			return
		}

		model.currentSize += int64(n)
	}
}
