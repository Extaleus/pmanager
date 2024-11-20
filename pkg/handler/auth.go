package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("signUp")
}

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIn")
}
