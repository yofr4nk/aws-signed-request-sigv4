package building_test

import (
	"fmt"
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildAuthenticatedHeadersShouldReturnStringWithAuthenticatedHeaders(t *testing.T) {
	expectedAuthenticatedHeaders := fmt.Sprintf("%s Credential=FakeAccessKey/2020/06/28/FakeRegion/s3/aws4_request, SignedHeaders=content-md5;host;x-amz-content-sha256;x-amz-date, Signature=FakeSignature", building.Algorithm)
	authenticatedHeaders := building.BuildAuthenticatedHeaders("FakeAccessKey", "2020/06/28", "FakeRegion", "FakeSignature")

	if authenticatedHeaders != expectedAuthenticatedHeaders {
		t.Errorf("Expected %v authenticated headers, but got %v", expectedAuthenticatedHeaders, authenticatedHeaders)
	}
}
