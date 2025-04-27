package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("calculator 1.0")
	fmt.Println("=======================")
	fmt.Println("Which operation do you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanln(&operation)

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	switch operation {
	case "add":
		fmt.Println("Result:", num1+num2)
	case "subtract":
		fmt.Println("Result:", num1-num2)
	case "multiply":
		fmt.Println("Result:", num1*num2)
	case "divide":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Println("Result:", num1/num2)
		}
	default:
		fmt.Println("Invalid operation.")
	}
}
