package building

import "github.com/yofr4nk/aws-signed-request-sigv4/hashing"

func BuildStringToSign(canonicalHeader string, amzDate string, dateStamp string, region string) string {
	var stringToSign string
	algorithm := "AWS4-HMAC-SHA256"

	stringToSign = algorithm + "\n"
	stringToSign += amzDate + "\n"
	stringToSign += dateStamp + "/" + region + "/s3/aws4_request\n"
	stringToSign += hashing.CreateSha256(canonicalHeader)

	return stringToSign
}
