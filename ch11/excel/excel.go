package excel

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
)

type ExcelRecord struct {
	Record map[string]string
}

func ReadXlsx(filename string) ([]ExcelRecord, error) {

	var allRecords = make([]ExcelRecord, 0)
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return allRecords, err
	} // 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		//fmt.Println("sheet name: ", sheet.Name) //遍历行读取
		//record := &ExcleRecord{make(map[string]string)}
		records := sheet.Rows[0].Cells
		colNum := len(records)
		recordNum := len(sheet.Rows)
		for i := 0; i < recordNum; i++ {
			record := &ExcelRecord{make(map[string]string)}
			for k := 0; k < colNum; k++ {
				cells := sheet.Rows[i].Cells
				record.Record[records[k].Value] = cells[k].Value
			}
			allRecords = append(allRecords, *record)
		}
	}
	return allRecords, nil
}
func ReadCsv(filename string, columns map[string]string) ([]ExcelRecord, error) {

	var allRecords []ExcelRecord
	h, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return allRecords, err
	}
	defer h.Close()
	r := transform.NewReader(h, simplifiedchinese.GBK.NewDecoder())
	reader := csv.NewReader(bufio.NewReader(r))
	if reader == nil {
		fmt.Println("NewReader return nil, file:")
		return allRecords, errors.New("NewReader return nil")
	}
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return allRecords, err
	}
	colNum := len(records[0])
	recordNum := len(records)
	for i := 0; i < recordNum; i++ {
		record := &ExcelRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[columns[records[0][k]]] = records[i][k]
		}
		allRecords = append(allRecords, *record)
	}
	return allRecords, nil
}
