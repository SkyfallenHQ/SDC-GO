package URLEncoder

import (
	"net/url"
)

func EncodeQS(s string) string {
	return url.QueryEscape(s)
}