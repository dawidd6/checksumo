package models_test

import (
	"testing"

	"github.com/dawidd6/checksumo/src/models"

	"github.com/stretchr/testify/assert"
)

func TestDetectType(t *testing.T) {
	m := models.NewMainModel()

	cases := map[string]string{
		"f3a306f40e4a313fb5a584d73b3dee8f":                                 "MD5",
		"443511f6bf12402c12503733059269a2e10dec602916c0a75263e5d990f6bb93": "SHA-256",
		"302c990c6d69575ff24c96566e5c7e26bf36908abb0cd546e22687c46fb07bf8dba595bf77a9d4fd9ab63e75c0437c133f35462fd41ea77f6f616140cd0e5e6a": "SHA-512",
	}

	for hashVal, expectedHashType := range cases {
		m.SetHash(hashVal)
		gotHashType := m.DetectType()
		assert.Equal(t, expectedHashType, gotHashType)
	}
}
