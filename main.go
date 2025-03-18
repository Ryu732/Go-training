package main

import (
	"Gin/controllers"
	"Gin/infra"
	"Gin/repositories"

	//"Gin/models"

	"Gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	//items := []models.Item{}

	//itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
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
