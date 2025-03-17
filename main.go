package main

import (
	"Gin/controllers"
	"Gin/models"
	"Gin/repositories"
	"Gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{
			ID:          1,
			Name:        "商品1",
			Price:       1000,
			Description: "説明1",
			Soldout:     false,
		},
		{
			ID:          2,
			Name:        "商品2",
			Price:       2000,
			Description: "説明2",
			Soldout:     true,
		},
		{
			ID:          3,
			Name:        "商品3",
			Price:       3000,
			Description: "説明3",
			Soldout:     false,
		},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemContoroller(itemService)

	router := gin.Default()

	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindById)
	router.POST("/items", itemController.Create)
	router.PUT("/items/:id", itemController.Update)
	router.DELETE("/items/:id", itemController.Delete)

	router.Run("localhost:8080")
}
