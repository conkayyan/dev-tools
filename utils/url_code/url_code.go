package url_code

import (
	"net/url"
)

func Encode(str string) string {
	return url.QueryEscape(str)
}
func Decode(str string) string {
	data, err := url.QueryUnescape(str)
	if err != nil {
		return err.Error()
	}
	return data
}
