package handlers

import (
	"log"
	"net/http"
	"products/internal/products/model/apperrors"

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
	ctx := c.Request.Context()
	p, err := h.ProductService.Get(ctx, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", p, err)
		e := apperrors.NewNotFound("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": p,
	})
}
