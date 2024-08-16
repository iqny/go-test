package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

func main() {
	/*f, err := excelize.OpenFile("C:\\Users\\Administrator\\Downloads\\26c76331ca96464c9be0fd1dabc0768b.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(f.GetSheetList())
	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}*/

	var (
		user     = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
		employee = Employee{Salary: 150000}
	)

	copier.Copy(&user, &employee)

	fmt.Printf("%#v \n", user)
}

type User struct {
	Name         string
	Role         string
	Age          int32
	EmployeeCode int64 `copier:"EmployeeNum"` // specify field name

	// Explicitly ignored in the destination struct.
	Salary int
}
type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	Name string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	Age int32 `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	Salary int `copier:"-"`

	DoubleAge  int32
	EmployeeId int64 `copier:"EmployeeNum"` // specify field name
	SuperRole  string
}
