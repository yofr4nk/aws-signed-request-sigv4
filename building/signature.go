package building

import "github.com/yofr4nk/aws-signed-request-sigv4/hashing"

func BuildSignature(signingKey string, stringToSign string) string {
	return hashing.CreateHmac(signingKey, stringToSign)
}
