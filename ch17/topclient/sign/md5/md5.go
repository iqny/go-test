package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func GenerateMd5(body string) string {
	w := md5.New()
	_, err := io.WriteString(w, body)
	if err != nil {
		return ""
	}
	return strings.ToUpper(fmt.Sprintf("%x", w.Sum(nil)))
}
