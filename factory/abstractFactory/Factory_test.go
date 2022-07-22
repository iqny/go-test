package abstractFactory

import "testing"

func TestCreateFactory(t *testing.T) {
	factory := CreateFactory("Mysql")
	score := factory.CreateScore()
	score.Insert()
	score.List()
	student := factory.CreateStudent()
	student.Insert()
	student.Update()
}
