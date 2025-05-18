package main

import "fmt"

func LinearSearch(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			fmt.Println("Key is found at index:", i)
			return i
		}
	}
	fmt.Println("Key not found")
	return -1
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	result := LinearSearch(numbers, 3)
	fmt.Println("Search result index:", result)
}
