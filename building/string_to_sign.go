package building

import "github.com/yofr4nk/aws-signed-request-sigv4/hashing"

func BuildStringToSign(canonicalHeader string, fullDate string, dateYMD string, region string) string {
	var stringToSign string

	stringToSign = "AWS4-HMAC-SHA256\n"
	stringToSign += fullDate + "\n"
	stringToSign += dateYMD + "/" + region + "/s3/aws4_request\n"
	stringToSign += hashing.CreateSha256(canonicalHeader)

	return stringToSign
}
