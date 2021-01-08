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
}

func New() *Model {
	return &Model{}
}

func (model *Model) DetectType(h string) string {
	switch len(h) {
	case md5.Size * 2:
		model.hasher = md5.New()
		return "MD5"
	case sha256.Size * 2:
		model.hasher = sha256.New()
		return "SHA-256"
	case sha512.Size * 2:
		model.hasher = sha512.New()
		return "SHA-512"
	default:
		model.hasher = nil
		return ""
	}
}

func (model *Model) ComputeHash(ctx context.Context, filePath string) (string, error) {
	// Open file
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	// Reset hasher and close the file on exit
	defer func() {
		model.hasher.Reset()
		file.Close()
	}()

	// Read file
	for {
		// Check if context were cancelled
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		// Define 32kB buffer
		buffer := make([]byte, 32*1024)

		// Read bytes
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		buffer = buffer[:n]

		// Write bytes to hasher
		n, err = model.hasher.Write(buffer)
		if err != nil {
			return "", err
		}
	}

	// Compute checksum, encode return it
	return hex.EncodeToString(model.hasher.Sum(nil)), nil
}
