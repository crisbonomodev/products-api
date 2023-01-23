package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockProductsService struct {
	mock.Mock
}

func (m *MockProductsService) Get(ctx context.Context, uid uuid.UUID) {

}
