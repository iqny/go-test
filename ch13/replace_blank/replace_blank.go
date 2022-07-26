package replace_blank

import (
	"strings"
	"unicode"
)

func replaceBlank(s string) (string, bool) {
	if len(s) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) != false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
