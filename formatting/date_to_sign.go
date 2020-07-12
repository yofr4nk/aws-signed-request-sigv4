package formatting

import "time"

const (
	StampLayout  string = "20060102"
	AMZLayout string = "20060102T150405Z0700"
)

type DatesFormatToSign struct {
	DateStamp string
	AMZDate string
}

func BuildDatesToSign(dateToParse time.Time) DatesFormatToSign {
	dateStamp := dateToParse.Format(StampLayout)
	amzDate := dateToParse.Format(AMZLayout)

	return DatesFormatToSign{
		DateStamp: dateStamp,
		AMZDate: amzDate,
	}
}
