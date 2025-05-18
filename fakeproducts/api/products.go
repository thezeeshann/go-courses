package api

import (
	"encoding/json"
	"net/http"

	"github.com/thezeeshann/fake-products/model"
)

const apiUrl = "https://api.escuelajs.co/api/v1/products"

func GetProducts() ([]model.Products, error) {
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var products []model.Products
	err = json.NewDecoder(res.Body).Decode(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
