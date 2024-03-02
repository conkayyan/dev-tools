package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func GenUUID(isHyphen bool) string {
	uuidWithHyphen := uuid.NewString()
	if !isHyphen {
		return strings.ReplaceAll(uuidWithHyphen, "-", "")
	}
	return uuidWithHyphen
}
