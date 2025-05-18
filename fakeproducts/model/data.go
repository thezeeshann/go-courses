package model

type Products struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Price       int32  `json:"price"`
	Description string `json:"description"`
}
