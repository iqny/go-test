package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	gob.Register([]map[string]interface{}{})
}
func main() {
	var skuSn = make([]int64, 0)
	var skuSnString string
	skuSn = []int64{1, 2, 3, 4}
	for _, sku := range skuSn {
		skuSnString = fmt.Sprintf("%s\"%s\",", skuSnString, strconv.FormatInt(sku, 10))
	}
	fmt.Println(strings.Trim(skuSnString, ","))
	var bb = []map[string]interface{}{{"abc": 123}}
	var mm = map[string]interface{}{"abc": bb}
	by, err := byteEncoder(mm)
	fmt.Println(err)
	fmt.Println(by)
}
func byteEncoder(s interface{}) ([]byte, error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
