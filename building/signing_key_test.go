package building_test

import (
	"encoding/hex"
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"testing"
)

func TestBuildSigningKeyShouldReturnSignatureKeyHashed(t *testing.T) {
	expectedSingingKey := "0411f58aca6b8f65c409bf63725b1ea28389b7aeae7d1826b30603977649f867"

	signingKey := building.BuildSigningKey("fakeKey", "20200605", "us-east-1")

	if hex.EncodeToString(signingKey) != expectedSingingKey {
		t.Errorf("Expected %v signingKey key, but got %v", expectedSingingKey, hex.EncodeToString(signingKey))
	}
}
