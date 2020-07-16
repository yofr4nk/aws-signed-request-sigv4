package sign

import (
	"fmt"
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"github.com/yofr4nk/aws-signed-request-sigv4/formatting"
	"time"
)

const XAmzContentSha256 string = "UNSIGNED-PAYLOAD"

type PayloadToSing struct {
	Date       time.Time
	KeyPath    string
	UploadId   string
	PartNumber string
	ContentMD5 string
	Method     string
}

type ConfigData struct {
	Bucket          string
	SecretAccessKey string
	AccessKeyId     string
	Region          string
}

type HeaderData struct {
	Host              string
	XAmzContentSha256 string
	XAmzDate          string
	ContentMD5        string
	Authorization     string
}

type SignatureResponse struct {
	Url    string
	Header HeaderData
}

func CalculateSignature(p PayloadToSing, cd ConfigData) SignatureResponse {
	dateFormats := formatting.BuildDatesToSign(p.Date)
	qp, keysOrder := buildQueryParams(p)

	paramsToSign := building.BuildPathToSign(qp, keysOrder, cd.Bucket)
	canonicalHeaders := building.BuildCanonicalHeader(building.CHParams{
		Method:      p.Method,
		KeyPath:     p.KeyPath,
		QueryString: paramsToSign.QueryString,
		Host:        paramsToSign.Host,
		AMZDate:     dateFormats.AMZDate,
		ContentMD5:  p.ContentMD5,
	})
	stringToSign := building.BuildStringToSign(canonicalHeaders, dateFormats.AMZDate, dateFormats.DateStamp, cd.Region)
	// Create the signing key to be used in BuildSignature function.
	signingKey := building.BuildSigningKey(cd.SecretAccessKey, dateFormats.DateStamp, cd.Region)
	// Sign the string_to_sign using the signing_key
	signature := building.BuildSignature(signingKey, stringToSign)
	authenticatedHeaders := building.BuildAuthenticatedHeaders(cd.AccessKeyId, dateFormats.DateStamp, cd.Region, signature)

	return SignatureResponse{
		Url: fmt.Sprintf("https://%s/%s?%s", paramsToSign.Host, p.KeyPath, paramsToSign.QueryString),
		Header: HeaderData{
			Host:              paramsToSign.Host,
			XAmzContentSha256: XAmzContentSha256,
			XAmzDate:          dateFormats.AMZDate,
			ContentMD5:        p.ContentMD5,
			Authorization:     authenticatedHeaders,
		},
	}
}

func buildQueryParams(p PayloadToSing) (map[string]string, []string) {
	qp := make(map[string]string)

	qp["partNumber"] = p.PartNumber
	qp["uploadId"] = p.UploadId

	// It's needed to keep the order iterating QueryParams
	keysOrder := []string{"partNumber", "uploadId"}

	return qp, keysOrder
}
