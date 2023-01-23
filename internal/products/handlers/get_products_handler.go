package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetProducts(c *gin.Context) {
	productId, exists := c.Params.Get("id")
	if !exists {
		log.Printf("Unable to extract product from request context for unknown reason: %v\n", c)
	}
	uid, err := uuid.Parse(productId)
	if err != nil {
		log.Printf("Unable to cast string to UUID: %v\n", c)
	}
	// gin.Context satisfies go's context.Context interface
	p, err := h.ProductService.Get(c, uid)

	if err != nil {
		log.Printf("Unable to find product: %v\n%v", uid, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": p,
	})
}
