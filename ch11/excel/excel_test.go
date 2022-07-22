package excel

import (
	"fmt"
	"testing"
)

func TestReadXlsx(t *testing.T) {
	record,err:=ReadXlsx("D:/juest/src/go-zh/ch11/excel/88.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println(record)
	record = record[2:]
	for _, dt := range record {
		fmt.Println(dt.Record["cat"])
		fmt.Println(dt.Record["goods_sn"])
		fmt.Println(dt.Record["spec"])
		fmt.Println(dt.Record["spec_value"])
		fmt.Println(dt.Record["sku_sn"])
		fmt.Println(dt.Record["price"])
	}
}
