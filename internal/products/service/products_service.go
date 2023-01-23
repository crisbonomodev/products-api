package service

import (
	"context"
	"products/internal/products/model"

	"github.com/google/uuid"
)

// ProductService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type ProductService struct {
	ProductRepository model.ProductRepository
}

// ProductConfig will hold repositories that will eventually be injected into this
// this service layer
type ProductConfig struct {
	ProductRepository model.ProductRepository
}

func NewProductService(c *ProductConfig) model.ProductService {
	return &ProductService{
		ProductRepository: c.ProductRepository,
	}
}

func (p *ProductService) Get(ctx context.Context, uid uuid.UUID) (*model.Product, error) {
	product := model.Product{
		UID: uid,
	}
	return &product, nil
	// u, err := p.ProductRepository.FindByID(ctx, uid)

	// return u, err
}
