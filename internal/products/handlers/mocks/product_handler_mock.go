package mocks

import (
	"context"
	"products/internal/products/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// Mock para model.ProductService
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Get(ctx context.Context, uid uuid.UUID) (*model.Product, error) {
	//argumentos que se pasaran a "return" en la funcion, cuando se llame con un uid o error
	ret := m.Called(ctx, uid)

	// Primer valor que pasarmos a return
	var r0 *model.Product
	//si pasamos un uid devolvemos este producto
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Product)
	}

	var r1 error
	//si no pasamos un uid devolvemos este error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockProductService) Create(ctx context.Context, product *model.Product) error {
	ret := m.Called(ctx, product)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
