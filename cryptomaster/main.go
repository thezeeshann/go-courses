package main

import (
	"fmt"

	"github.com/thezeeshann/crypto/api"
)

func main() {

	rate, err := api.GetRate("BTC")
	fmt.Print(rate, err)

}

// create one project and call fake api and show the data in the terminal -> porject
