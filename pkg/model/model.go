package model

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
)

type Model struct {
	File   string
	Hash   string
	hasher hash.Hash
}

func New() *Model {
	return &Model{}
}

func (model *Model) DetectType(h string) (string, error) {
	switch len(h) {
	case md5.Size * 2:
		model.hasher = md5.New()
		return "MD5", nil
	case sha256.Size * 2:
		model.hasher = sha256.New()
		return "SHA-256", nil
	case sha512.Size * 2:
		model.hasher = sha512.New()
		return "SHA-512", nil
	default:
		model.hasher = nil
		return "", errors.New("hash type unrecognized")
	}
}

func (model *Model) ComputeHash(ctx context.Context, filePath string) (string, error) {
	defer model.hasher.Reset()

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	for {
		select {
		case <-ctx.Done():
			return "", err
		default:
		}

		buffer := make([]byte, 32768)

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
