package Controllers

import (
	"example.com/api_prac/Models"
	"example.com/api_prac/initializers"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []Models.Product
	initializers.DB.Find(&products)

	c.JSON(200, gin.H{
		"Products": products,
	})
}

func CreateProduct(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Price       float32
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	product := Models.Product{Title: body.Title, Description: body.Description, Price: body.Price}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"Message": "Error Inserting mto database",
		})
	}

	c.JSON(201, gin.H{
		"Message": product,
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product Models.Product
	initializers.DB.First(&product, id)

	c.JSON(200, gin.H{
		"Product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title       string
		Description string
		Price       float32
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var product Models.Product
	initializers.DB.First(&product, id)

	initializers.DB.Model(&product).Updates(Models.Product{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})

	c.JSON(200, gin.H{
		"Product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product Models.Product
	initializers.DB.Delete(&product, id)

	c.JSON(200, gin.H{
		"Message": "Delete successful",
	})
}
