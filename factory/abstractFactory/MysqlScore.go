package abstractFactory

import "fmt"

type MysqlScore struct{}

func (m *MysqlScore) Insert() bool {
	fmt.Println("MysqlScore Insert")
	return true
}

func (m *MysqlScore) List() map[string]string {
	fmt.Println("MysqlScore List")
	return nil
}
