package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildSignatureShouldReturnAHashedSignature(t *testing.T) {
	signingKey := "fakeSigningKey"
	stringToSign := "fakeStringToSign"
	expectedSignature := "2995967988a0e18b6f77e5319f0e4869ef3d4d3bdc0ede53652abb46d1fbafad"

	signature := building.BuildSignature(signingKey, stringToSign)

	if signature != expectedSignature {
		t.Errorf("Expected %v signature, but got %v", expectedSignature, signature)
	}
}
