package abstractFactory

type Factory interface {
	CreateStudent() IStudent
	CreateScore() IScore
}

func CreateFactory(mod string) Factory {
	switch mod {
	case "Mysql":
		return new(MysqlFactory)
	case "Access":
		return new(AccessFactory)
	}
	return nil
}
