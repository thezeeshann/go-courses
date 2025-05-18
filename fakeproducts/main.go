package main

import (
	"fmt"

	"github.com/thezeeshann/fake-products/api"
)

func main() {

	products, err := api.GetProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range products {
		fmt.Printf("ID: %d, Title: %s, Price: %d\n", p.ID, p.Title, p.Price)
	}

}
