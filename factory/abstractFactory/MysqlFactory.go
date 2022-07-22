package abstractFactory

type MysqlFactory struct{}

func (m *MysqlFactory) CreateStudent() IStudent {
	return new(MysqlStudent)
}

func (m *MysqlFactory) CreateScore() IScore {
	return new(MysqlScore)
}
