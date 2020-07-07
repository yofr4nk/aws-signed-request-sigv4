package hashing_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/hashing"
	"testing"
)

func TestCreateSha256ShouldBuildAHashBasedOnSha256Standard(t *testing.T) {
	h := "63c232b252235370cccec00bb7c102e241ab3822047be1d50ff00ba6a0da3f4c"

	textHashed := hashing.CreateSha256("fake message to hash")
	if textHashed != h {
		t.Errorf("Expected %v text hashed based on Sha256 standard, but got %v", h, textHashed)
	}
}

func TestCreateHmacShouldBuildHashBasedOnKeyedHashMessageAuthenticationCode(t *testing.T) {
	h := "cbc2f42e095370cf0e36a22fb38cc2da8cf4afcad8e697ef402e8da13b4ca811"
	key := "MyTestKey"

	textHashed := hashing.CreateHmac(key, "fake message to hash")
	if textHashed != h {
		t.Errorf("Expected %v text hashed based on Hmac, but got %v", h, textHashed)
	}
}
