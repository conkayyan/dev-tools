package uuid

import (
	"strings"

	u "github.com/google/uuid"
)

func GenUUID(isHyphen bool) string {
	uuidWithHyphen := u.NewString()
	if !isHyphen {
		return strings.ReplaceAll(uuidWithHyphen, "-", "")
	}
	return uuidWithHyphen
}
