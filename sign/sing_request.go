package sign

const XAmzContentSha256 string = "UNSIGNED-PAYLOAD"

type PayloadToSing struct {
	keyPath    string
	uploadId   string
	partNumber string
	contentMD5 string
}

type ConfigData struct {
	bucket          string
	secretAccessKey string
	accessKeyId     string
}

type HeaderData struct {
	Host              string
	XAmzContentSha256 string
	XAmzDate          string
	ContentMD5        string
	Authorization     string
}

type SignatureResponse struct {
	Url    string
	Header HeaderData
}

func CalculateSignature(p PayloadToSing, cd ConfigData) SignatureResponse {
	return SignatureResponse{}
}
