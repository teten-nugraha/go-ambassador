package models

type Product struct {
	Id          uint	`json:"id"`
	Title       string	`json:"title"`
	Description string	`json:"description"`
	Image       string	`json:"Image"`
	Price       float64	`json:"price"`
}
