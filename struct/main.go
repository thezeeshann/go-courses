package main

import (
	"fmt"
	"struct/model"
)

func main() {
	zee := model.Instructor{FirstName: "zeeshan", LastName: "khan", Score: 22}
	kayil := model.NewInstructor("kayile", "desua")
	goCourse := model.Course{Id: 2, Name: "Go fundametnals"}
	swiftWs := model.Workshop{}

	fmt.Printf("%v", goCourse)
	fmt.Printf("%v", swiftWs)
	println(zee.Print())
	println(kayil.Print())

}
