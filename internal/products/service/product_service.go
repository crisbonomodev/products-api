package service

import (
	"products/internal/products/model"
)

// En este service se inyectan las dependencias que vamos a necesitar luego.
type productService struct {
	ProductRepository model.ProductRepository
}

// ProductConfig will hold repositories that will eventually be injected into this
// this service layer
type ProductConfig struct {
	ProductRepository model.ProductRepository
}

func NewProductService(c *ProductConfig) model.ProductService {
	return &productService{
		ProductRepository: c.ProductRepository,
	}
}
