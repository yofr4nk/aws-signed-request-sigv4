package building

import "github.com/yofr4nk/aws-signed-request-sigv4/hashing"

const Algorithm string = "AWS4-HMAC-SHA256"

func BuildStringToSign(canonicalHeader string, amzDate string, dateStamp string, region string) string {
	var stringToSign string

	stringToSign = Algorithm + "\n"
	stringToSign += amzDate + "\n"
	stringToSign += dateStamp + "/" + region + "/s3/aws4_request\n"
	stringToSign += hashing.CreateSha256(canonicalHeader)

	return stringToSign
}
