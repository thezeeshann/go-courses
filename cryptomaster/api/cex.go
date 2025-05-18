package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/thezeeshann/crypto/model"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) { // pointer can be nil
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(bodyBytes)
		fmt.Print(json)
	} else {
		return nil, fmt.Errorf("status code recived %v ", res.StatusCode) // creating own error
	}
	rate := model.Rate{Currency: currency, Price: 20}
	return &rate, nil // becase rate is pointer

}
