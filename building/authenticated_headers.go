package building

import "fmt"

func BuildAuthenticatedHeaders(accessKeyId string, dateStamp string, region string, signature string) string {
	authHeaders := fmt.Sprintf("%s Credential=%s/%s/%s/s3/aws4_request, ", Algorithm, accessKeyId, dateStamp, region)

	authHeaders += fmt.Sprintf("SignedHeaders=content-md5;host;x-amz-content-sha256;x-amz-date, Signature=%s", signature)

	return authHeaders
}
