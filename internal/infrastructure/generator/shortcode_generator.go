package generator

import (
	"crypto/rand"
	"encoding/base64"
)

type ShortCodeGenerator struct{}

func NewShortCodeGenerator() *ShortCodeGenerator {
	return &ShortCodeGenerator{}
}

func (g *ShortCodeGenerator) Generate() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}
