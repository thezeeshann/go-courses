package main

import "fmt"

func main() {
	numbers := [5]int{3, 2, 4, 7, 5}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
