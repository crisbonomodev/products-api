package handlers

import (
	"log"
	"net/http"
	"products/internal/products/model"
	"products/internal/products/model/apperrors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	productId, exists := c.Params.Get("id")
	if !exists {
		log.Printf("Unable to extract product from request context for unknown reason: %v\n", c)
	}
	uid, err := uuid.Parse(productId)
	if err != nil {
		log.Printf("Unable to cast string to UUID: %v\n", c)
	}
	product := &model.Product{
		UID: uid,
	}
	ctx := c.Request.Context()
	errCreation := h.ProductService.Create(ctx, product)

	if errCreation != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}
