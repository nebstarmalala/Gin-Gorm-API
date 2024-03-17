package Models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Description string
	Price       float32
}
