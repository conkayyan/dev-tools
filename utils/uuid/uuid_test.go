package uuid_test

import (
	"dev-tools/utils/uuid"
	"testing"
)

func TestGenUUID(t *testing.T) {
	t.Log(uuid.GenUUID(true))
}

func TestGenUUIDNoHyphen(t *testing.T) {
	t.Log(uuid.GenUUID(false))
}
