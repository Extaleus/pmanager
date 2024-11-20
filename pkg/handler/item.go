package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	fmt.Println("createItem")
}

func (h *Handler) getAllItems(c *gin.Context) {
	fmt.Println("getAllItems")
}

func (h *Handler) getItemById(c *gin.Context) {
	fmt.Println("getItemById")
}

func (h *Handler) updateItem(c *gin.Context) {
	fmt.Println("updateItem")
}

func (h *Handler) deleteItem(c *gin.Context) {
	fmt.Println("deleteItem")
}
