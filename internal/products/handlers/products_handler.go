package handlers

import (
	"products/internal/products/model"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	ProductService model.ProductService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R              *gin.Engine
	ProductService model.ProductService
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		ProductService: c.ProductService,
	}

	v1 := c.R.Group("/api/v1")
	{
		v1.GET("/products/:id", h.GetProducts)
		v1.POST("/products", h.CreateProduct)
	}
}
