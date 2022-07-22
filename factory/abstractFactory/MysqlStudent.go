package abstractFactory

import "fmt"

type MysqlStudent struct{}

func (m *MysqlStudent) Insert() bool {
	fmt.Println("MysqlStudent Insert")
	return true
}

func (m *MysqlStudent) Update() bool {
	fmt.Println("MysqlStudent Update")
	return true
}
