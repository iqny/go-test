package abstractFactory

type AccessFactory struct{}

func (a *AccessFactory) CreateStudent() IStudent {
	return new(AccessStudent)
}

func (a *AccessFactory) CreateScore() IScore {
	return new(AccessScore)
}
