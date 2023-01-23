package model

import (
	"context"

	"github.com/google/uuid"
)

// ProductService defines methods the handler layer expects
// any service it interacts with to implement
type ProductService interface {
	Get(ctx context.Context, uid uuid.UUID) (*Product, error)
	// Update(ctx context.Context, uid uuid.UUID) (*Product, error)
	// Delete(ctx context.Context, uid uuid.UUID) (*Product, error)
	// Create(ctx context.Context, uid *Product) (*Product, error)
}

// ProductRepository defines methods the service layer expects
// any repository it interacts with to implement
type ProductRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*Product, error)
}
