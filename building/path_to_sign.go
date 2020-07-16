package building

import "fmt"

type paramsToSign struct {
	QueryString string
	Host        string
}

type queryParams map[string]string

func BuildPathToSign(qp queryParams, keyParams []string, bucket string) paramsToSign {
	return paramsToSign{
		QueryString: buildQueryStringToSign(qp, keyParams),
		Host:        fmt.Sprintf("%s.s3.amazonaws.com", bucket),
	}
}

func buildQueryStringToSign(qp queryParams, keyParams []string) string {
	var qs string
	paramsLength := len(qp)
	n := 0

	for n < paramsLength {
		var concat string
		kp := keyParams[n]

		if len(qs) > 0 {
			concat = "&"
		}
		if kp != "" {
			qs += fmt.Sprintf("%s%s=%s", concat, kp, qp[kp])
		}

		n++
	}

	return qs
}
