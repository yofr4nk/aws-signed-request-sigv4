package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildHostS3ToSign_ShouldBuildTheS3HostFromBucketPassed(t *testing.T) {
	bucket := "myFakeBucket"
	pathExpected := "myFakeBucket.s3.amazonaws.com"

	hostToSign := building.BuildS3Host(bucket)
	if hostToSign != pathExpected {
		t.Errorf("Expected %v bucket path, but got %v", pathExpected, hostToSign)
	}
}
