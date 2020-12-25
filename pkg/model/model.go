package model

import (
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
	case sha256.BlockSize:
		model.hasher = sha256.New()
		return "SHA-256", nil
	case sha512.BlockSize:
		model.hasher = sha512.New()
		return "SHA-512", nil
	default:
		model.hasher = nil
		return "", errors.New("hash type unrecognized")
	}
}

func (model *Model) ComputeHash(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	model.hasher.Reset()

	_, err = io.Copy(model.hasher, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(model.hasher.Sum(nil)), nil
}
