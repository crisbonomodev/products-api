package mocks

import (
	"context"
	"products/internal/products/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.Product, error) {
	ret := m.Called(ctx, uid)

	var r0 *model.Product
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Product)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	ret := m.Called(ctx, product)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
