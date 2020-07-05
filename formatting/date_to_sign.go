package formatting

import "time"

const (
	YMDLayout  string = "20060102"
	FullLayout string = "20060102T150405Z0700"
)

type DatesFormatToSign struct {
	YMD string
	UTC string
}

func BuildDatesToSign(dateToParse time.Time) DatesFormatToSign {
	ymd := dateToParse.Format(YMDLayout)
	nsd := dateToParse.Format(FullLayout)

	return DatesFormatToSign{
		YMD: ymd,
		UTC: nsd,
	}
}
