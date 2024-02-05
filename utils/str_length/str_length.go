package str_length

import (
	"unicode/utf8"
)

func Count(txt string) int {
	return utf8.RuneCountInString(txt)
}
