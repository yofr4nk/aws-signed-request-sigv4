package building

import "fmt"

func BuildS3Host(bucket string) string {
	return fmt.Sprintf("%s.s3.amazonaws.com", bucket)
}
