package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"strings"
	"testing"
)

func TestBuildStringToSignShouldReturnStringWithDataToSign(t *testing.T) {
	canonicalHeader := "POST\nfakePath/fake-image.png\nuploadID=FakeID&partNumber=1\ncontent-md5:144c9defac04969c7bfad8efaa8ea194\nhost:myFakeBucket.s3.amazonaws.com\nx-amz-content-sha256:UNSIGNED-PAYLOAD\nx-amz-date:20200605T122800Z\n\ncontent-md5;host;x-amz-content-sha256;x-amz-date\nUNSIGNED-PAYLOAD"
	cnHashedExpected := "83256ca37cd10f0543bf99b2bb7351e4e430602bbf98b7677db58225a09cee65"
	fullDate := "20200605T122800Z"
	dateYMD := "20200605"
	region := "us-east-1"

	stringToSign := building.BuildStringToSign(canonicalHeader, fullDate, dateYMD, region)

	if strings.Contains(stringToSign, fullDate+"\n") == false {
		t.Errorf("Expected %v header, but got %v", fullDate, stringToSign)
	}

	if strings.Contains(stringToSign, dateYMD+"/"+region+"/s3/aws4_request\n") == false {
		t.Errorf("Expected %v header, but got %v", dateYMD+"/"+region+"/s3/aws4_request\n", stringToSign)
	}

	if strings.Contains(stringToSign, "AWS4-HMAC-SHA256\n") == false {
		t.Errorf("Expected %v header, but got %v", "AWS4-HMAC-SHA256\n", stringToSign)
	}

	if strings.Contains(stringToSign, cnHashedExpected) == false {
		t.Errorf("Expected %v header, but got %v", cnHashedExpected, stringToSign)
	}
}
