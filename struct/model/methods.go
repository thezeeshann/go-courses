package model

import "fmt"

// factories functions
func NewInstructor(name string, lastName string) Instructor {
	return Instructor{FirstName: name, LastName: lastName}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v, %d", i.FirstName, i.LastName, i.Score)
}
