package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildPathToSign_ShouldBuildTheS3HostFromBucketPassed(t *testing.T) {
	var queryParams map[string]string
	bucket := "myFakeBucket"
	pathExpected := "myFakeBucket.s3.amazonaws.com"
	var keyParams []string

	pathToSign := building.BuildPathToSign(queryParams, keyParams, bucket)
	if pathToSign.Path != pathExpected {
		t.Errorf("Expected %v bucket path, but got %v", pathExpected, pathToSign.Path)
	}

	if pathToSign.QueryString != "" {
		t.Errorf("Expected empty queryString, but got %v", pathToSign.QueryString)
	}
}

func TestBuildPathToSign_ShouldBuildTheQuery(t *testing.T) {
	queryParams := map[string]string{
		"uploadId":   "myTestID",
		"partNumber": "1",
	}
	keyParams := []string{"uploadId", "partNumber"}
	bucket := "myFakeBucket"
	pathExpected := "myFakeBucket.s3.amazonaws.com"
	qsExpected := "uploadId=myTestID&partNumber=1"

	pathToSign := building.BuildPathToSign(queryParams, keyParams, bucket)
	if pathToSign.Path != pathExpected {
		t.Errorf("Expected %v bucket path, but got %v", pathExpected, pathToSign.Path)
	}

	if pathToSign.QueryString != qsExpected {
		t.Errorf("Expected %v queryString, but got %v", qsExpected, pathToSign.QueryString)
	}
}
