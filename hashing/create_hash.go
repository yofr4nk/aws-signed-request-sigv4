package hashing

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CreateSha256(t string) string {
	h := sha256.Sum256([]byte(t))
	hexData := hex.EncodeToString(h[:])

	return hexData
}

func CreateHmac(key []byte, t string) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(t))

	return h.Sum(nil)
}
