package unique_string

import (
	"strings"
)


func isUniqueString(s string) (b bool) {
	if strings.Count(s,"")>3000{
		return
	}
	for _,v:=range s{
		if v > 127 {
			return
		}
		if strings.Count(s, string(v)) > 1 {
			return
		}
	}
	return true
}
