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

func CreateHmac(key string, t string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(t))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
