package building_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/building"
	"strings"
	"testing"
)

func TestBuildCanonicalHeaderShouldReturnHeaderToSignInOnlyString(t *testing.T) {
	chParams := building.CHParams{
		Method:      "POST",
		KeyPath:     "fakePath/fake-image.png",
		QueryString: "uploadID=FakeID&partNumber=1",
		Host:        "myFakeBucket.s3.amazonaws.com",
		AMZDate:    "20200605T122800Z",
		ContentMD5:  "144c9defac04969c7bfad8efaa8ea194",
	}

	canonicalHeader := building.BuildCanonicalHeader(chParams)

	if strings.Contains(canonicalHeader, chParams.Method) == false {
		t.Errorf("Expected %v header, but got %v", chParams.Method, canonicalHeader)
	}

	if strings.Contains(canonicalHeader, chParams.KeyPath) == false {
		t.Errorf("Expected %v header, but got %v", chParams.KeyPath, canonicalHeader)
	}

	if strings.Contains(canonicalHeader, chParams.QueryString) == false {
		t.Errorf("Expected %v header, but got %v", chParams.QueryString, canonicalHeader)
	}

	if strings.Contains(canonicalHeader, chParams.Host) == false {
		t.Errorf("Expected %v header, but got %v", chParams.Host, canonicalHeader)
	}

	if strings.Contains(canonicalHeader, chParams.AMZDate+"\n\n") == false {
		t.Errorf("Expected %v header, but got %v", chParams.AMZDate, canonicalHeader)
	}

	if strings.Contains(canonicalHeader, chParams.ContentMD5) == false {
		t.Errorf("Expected %v header, but got %v", chParams.ContentMD5, canonicalHeader)
	}
}
