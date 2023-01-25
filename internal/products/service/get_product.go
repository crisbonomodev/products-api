package service

import (
	"context"
	"products/internal/products/model"

	"github.com/google/uuid"
)

func (p *productService) Get(ctx context.Context, uid uuid.UUID) (*model.Product, error) {
	u, err := p.ProductRepository.FindByID(ctx, uid)

	return u, err
}
