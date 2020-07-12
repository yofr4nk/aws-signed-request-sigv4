package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildSignatureKeyShouldReturnSignatureKeyHashed(t *testing.T) {
	expectedSingingKey := "0c637c0a1e2b1282b91969a0a805f36b55c7e9a6937abdc578c07d2786e2bf6e"

	signatureKey := building.BuildSignatureKey("fakeKey", "20200605", "us-east-1")

	if signatureKey != expectedSingingKey {
		t.Errorf("Expected %v signature key, but got %v", expectedSingingKey, signatureKey)
	}
}
