package main

import (
	"example.com/api_prac/Models"
	"example.com/api_prac/initializers"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&Models.Product{})
	if err != nil {
		return
	}
}
