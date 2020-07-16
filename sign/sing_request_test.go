package sign_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/sign"
	"testing"
	"time"
)

func TestCalculateSignature(t *testing.T) {
	l, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error(err)
	}
	d := time.Date(2020, 07, 16, 18, 04, 15, 00, l)

	dataToSign := sign.PayloadToSing{
		Date:        d,
		KeyPath:     "fakePath/fakeFile.png",
		QueryString: "partNumber=1&uploadId=fakeUploadId",
		ContentMD5:  "144c9defac04969c7bfad8efaa8ea194",
		Method:      "PUT",
	}

	dataConfig := sign.ConfigData{
		Bucket:          "fakeBucket",
		SecretAccessKey: "fakeSecretAccessKey",
		AccessKeyId:     "fakeAccessKey",
		Region:          "us-east-1",
	}

	expectedSignature := sign.SignatureResponse{
		Url: "https://fakeBucket.s3.amazonaws.com/fakePath/fakeFile.png?partNumber=1&uploadId=fakeUploadId",
		Header: sign.HeaderData{
			Host:              "fakeBucket.s3.amazonaws.com",
			XAmzContentSha256: "UNSIGNED-PAYLOAD",
			XAmzDate:          "20200716T180415Z",
			ContentMD5:        "144c9defac04969c7bfad8efaa8ea194",
			Authorization:     "AWS4-HMAC-SHA256 Credential=fakeAccessKey/20200716/us-east-1/s3/aws4_request, SignedHeaders=content-md5;host;x-amz-content-sha256;x-amz-date, Signature=21e7bcad31cedcc2577b842f661781f3033bd5a045d557972d9166ec9fc4b483",
		},
	}

	signature := sign.CalculateSignature(dataToSign, dataConfig)

	if signature.Url != expectedSignature.Url {
		t.Errorf("Expected Url: %v, but got %v", expectedSignature.Url, signature.Url)
	}

	if signature.Header.XAmzDate != expectedSignature.Header.XAmzDate {
		t.Errorf("Expected XAmzDate: %v, but got %v", expectedSignature.Header.XAmzDate, signature.Header.XAmzDate)
	}

	if signature.Header.Host != expectedSignature.Header.Host {
		t.Errorf("Expected Host: %v, but got %v", expectedSignature.Header.Host, signature.Header.Host)
	}

	if signature.Header.XAmzContentSha256 != expectedSignature.Header.XAmzContentSha256 {
		t.Errorf("Expected XAmzContentSha256: %v, but got %v", expectedSignature.Header.XAmzContentSha256, signature.Header.XAmzContentSha256)
	}

	if signature.Header.ContentMD5 != expectedSignature.Header.ContentMD5 {
		t.Errorf("Expected ContentMD5: %v, but got %v", expectedSignature.Header.ContentMD5, signature.Header.ContentMD5)
	}

	if signature.Header.Authorization != expectedSignature.Header.Authorization {
		t.Errorf("Expected Authorization: %v, but got %v", expectedSignature.Header.Authorization, signature.Header.Authorization)
	}
}
