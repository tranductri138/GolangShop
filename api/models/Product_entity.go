package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	IsSale      bool    `json:"is_sale"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
