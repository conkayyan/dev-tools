package base64

import (
	"encoding/base64"
)

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
