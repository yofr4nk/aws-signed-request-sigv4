package building

import "github.com/yofr4nk/aws-signed-request-sigv4/hashing"

func BuildSigningKey(secretAccessKey string, dateStamp string, region string) string {
	keyDate := hashing.CreateHmac("AWS4"+secretAccessKey, dateStamp)
	dateRegionKey := hashing.CreateHmac(keyDate, region)
	dateRegionServiceKey := hashing.CreateHmac(dateRegionKey, "s3")
	signingKey := hashing.CreateHmac(dateRegionServiceKey, "aws4_request")

	return signingKey
}
