package building

import (
	"encoding/hex"
	"github.com/yofr4nk/aws-signed-request-sigv4/hashing"
)

func BuildSignature(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(hashing.CreateHmac(signingKey, stringToSign))
}
