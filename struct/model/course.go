package model

type Course struct {
	Id     int
	Name   string
	Slug   string
	Legacy bool
}

func (c Course) SingUp() bool {
	return true
}
