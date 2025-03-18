package main

import (
	"Gin/infra"
	"Gin/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("faild to migrate database")
	}
}
