package sign

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strings"
)
// CreateSign 生成验证码
func CreateSign(data map[string]string) string {
	sp := make(sysParamsSort, len(data))
	i := 0
	for k, v := range data {
		sp[i] = paramSort{k, v}
		i++
	}
	sort.Sort(sp)
	var query string
	for _, p := range sp {
		query += strings.Join([]string{p.Key,"=",p.Value,"&"},"")
	}
	return generateMd5(strings.Trim(query, "&"))
}

type paramSort struct {
	Key   string
	Value string
}

type sysParamsSort []paramSort

func (s sysParamsSort) Len() int           { return len(s) }
func (s sysParamsSort) Less(i, j int) bool { return s[i].Value < s[j].Value }
func (s sysParamsSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// generateMd5 生成md5
func generateMd5(body string) string {
	w := md5.New()
	_, err := io.WriteString(w, body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", w.Sum(nil))
}
