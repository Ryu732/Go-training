package controllers

import (
	"Gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemContoroller(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected erro"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
