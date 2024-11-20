package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	fmt.Println("createList")
}

func (h *Handler) getAllLists(c *gin.Context) {
	fmt.Println("getAllLists")
}

func (h *Handler) getListById(c *gin.Context) {
	fmt.Println("getListById")
}

func (h *Handler) updateList(c *gin.Context) {
	fmt.Println("updateList")
}

func (h *Handler) deleteList(c *gin.Context) {
	fmt.Println("deleteList")
}
