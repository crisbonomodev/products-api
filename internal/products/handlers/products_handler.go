package handlers

import (
	"products/internal/products/model"

	"github.com/gin-gonic/gin"
)

// Handler struct contiene los servicios necesarios para que el Handler funcione
type Handler struct {
	ProductService model.ProductService
}

// Config contiene las interfaces de los servicios que van a ser injectados en el handler cuando se inicialice
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
	v1.GET("/products/:id", h.GetProducts)
	v1.POST("/products", h.CreateProduct)
}
