package main

import "fmt"

// var url = "https://www.example.com"

func calculatePercentage(price float32) (float32, float32) {
	discount := price * 0.1
	tax := price * 0.2
	return discount, tax
}

func calculatePercentageWithName(price float32) (stateText float32, cityText float32) {
	stateText = price * 0.1
	cityText = price * 0.2
	return

}

func birthDay(age *int) {
	*age = *age + 1
}

func main() {

	age := 10
	birthDay(&age)
	fmt.Println("Age after birthday:", age)

	// stateText, _ := calculatePercentage(1000)
	// _, t := calculatePercentageWithName(1000)
	// fmt.Println(stateText, t)

	// var message string = "Hello, World!"
	// var message2 = "Hello, World!"
	// message3 := "Hello, World!"
	// prince := 12
	// println(message, message2, message3, prince, url)

	printData()

}
