package models

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

type MainModel struct {
	hasher hash.Hash

	providedHash string
	computedHash string
	filePath     string
	hashType     string

	totalBytes int64
	readBytes  int64

	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewMainModel() *MainModel {
	return &MainModel{}
}

func (model *MainModel) SetFile(f string) {
	model.filePath = f
}

func (model *MainModel) SetHash(h string) {
	model.providedHash = h
}

func (model *MainModel) GetProgress() float64 {
	return float64(model.readBytes) / float64(model.totalBytes)
}

func (model *MainModel) DetectType() string {
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

func (model *MainModel) IsReady() bool {
	return model.hashType != "" && model.filePath != "" && model.providedHash != ""
}

func (model *MainModel) PrepareHashing() {
	model.ctx, model.cancelFunc = context.WithCancel(context.Background())
}

func (model *MainModel) StopHashing() {
	model.cancelFunc()
}

func (model *MainModel) StartHashing() (bool, error) {
	// Cancel context on exit
	defer model.cancelFunc()

	// Open file
	file, err := os.OpenFile(model.filePath, os.O_RDONLY, 0666)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Get file info
	stat, err := file.Stat()
	if err != nil {
		return false, err
	}

	// Reset hasher buffer
	model.hasher.Reset()

	// Zero bytes
	model.totalBytes = stat.Size()
	model.readBytes = 0

	// Read file
	for {
		// Check if context were cancelled
		select {
		case <-model.ctx.Done():
			return false, model.ctx.Err()
		default:
		}

		// Define buffer
		buffer := make([]byte, 32*1024)

		// Read bytes
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
		buffer = buffer[:n]

		// Write bytes to hasher
		n, err = model.hasher.Write(buffer)
		if err != nil {
			return false, err
		}

		// Append read bytes
		model.readBytes += int64(n)
	}

	// Set computed hash
	model.computedHash = hex.EncodeToString(model.hasher.Sum(nil))

	// Return the result of hashes comparison
	return model.computedHash == model.providedHash, nil
}
