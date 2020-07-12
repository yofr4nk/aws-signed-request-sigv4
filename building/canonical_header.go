package building

type CHParams struct {
	Method      string
	KeyPath     string
	QueryString string
	Host        string
	FullDate    string
	ContentMD5  string
}

func BuildCanonicalHeader(chParams CHParams) string {
	var canonicalHeader string

	canonicalHeader = chParams.Method + "\n"
	canonicalHeader += chParams.KeyPath + "\n"
	canonicalHeader += chParams.QueryString + "\n"
	canonicalHeader += "content-md5:" + chParams.ContentMD5 + "\n"
	canonicalHeader += "host:" + chParams.Host + "\n"
	canonicalHeader += "x-amz-content-sha256:UNSIGNED-PAYLOAD\n"
	canonicalHeader += "x-amz-date:" + chParams.FullDate + "\n\n"
	canonicalHeader += "content-md5;host;x-amz-content-sha256;x-amz-date\n"
	canonicalHeader += "UNSIGNED-PAYLOAD"

	return canonicalHeader
}
