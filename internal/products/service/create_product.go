package service

import (
	"context"
	"products/internal/products/model"
)

func (p *productService) Create(ctx context.Context, product *model.Product) error {
	if err := p.ProductRepository.CreateProduct(ctx, product); err != nil {
		return err
	}

	return nil
}
