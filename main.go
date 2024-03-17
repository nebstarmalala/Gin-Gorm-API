package main

import (
	"example.com/api_prac/Controllers"
	"example.com/api_prac/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/products", Controllers.GetProducts)
	r.POST("/products", Controllers.CreateProduct)
	r.GET("/products/:id", Controllers.GetProduct)
	r.PUT("/products/:id", Controllers.UpdateProduct)
	r.DELETE("/products/:id", Controllers.DeleteProduct)

	err := r.Run()
	if err != nil {
		return
	}
}
