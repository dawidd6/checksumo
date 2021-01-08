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
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer func() {
		model.hasher.Reset()
		file.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		buffer := make([]byte, 32*1024)

		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		buffer = buffer[:n]

		n, err = model.hasher.Write(buffer)
		if err != nil {
			return "", err
		}
	}

	return hex.EncodeToString(model.hasher.Sum(nil)), nil
}
