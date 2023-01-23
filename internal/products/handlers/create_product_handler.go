package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}
