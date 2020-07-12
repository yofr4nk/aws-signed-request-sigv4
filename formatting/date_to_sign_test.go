package formatting_test

import (
	"github.com/yofr4nk/aws-signed-request-sigv4/formatting"
	"testing"
	"time"
)

func TestBuildDatesToSignShouldParseDateToYMDLayout(t *testing.T) {
	l, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error(err)
	}
	d := time.Date(2020, 06, 05, 12, 28, 00, 00, l)

	dts := formatting.BuildDatesToSign(d)
	if dts.DateStamp != "20200605" {
		t.Errorf("Expected %v with YMD format to sign, but got %v", "20200605", dts.DateStamp)
	}
}

func TestBuildDatesToSignShouldParseDateToFullUTCLayout(t *testing.T) {
	l, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error(err)
	}
	d := time.Date(2020, 06, 05, 12, 28, 00, 00, l)

	dts := formatting.BuildDatesToSign(d)
	if dts.AMZDate != "20200605T122800Z" {
		t.Errorf("Expected %v with AMZDate format to sign, but got %v", "20200605T122800Z", dts.AMZDate)
	}
}
