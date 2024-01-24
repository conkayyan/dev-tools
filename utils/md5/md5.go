package md5

import (
	utilMd5 "crypto/md5"
	"encoding/hex"
)

func Encode(str string) string {
	md5New := utilMd5.New()
	md5New.Write([]byte(str))
	return hex.EncodeToString(md5New.Sum(nil))
}
