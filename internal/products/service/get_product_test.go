package service

import (
	"context"
	"products/internal/products/model"
	"products/internal/products/model/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		//generamos un UUID random
		uid, _ := uuid.NewRandom()
		//mockeamos la respuesta del FindByID
		mockProductResp := &model.Product{
			UID: uid,
		}
		//mockeamos el repository
		MockProductRepository := new(mocks.MockProductRepository)
		mockProductService := NewProductService(&ProductConfig{
			ProductRepository: MockProductRepository,
		})
		MockProductRepository.On("FindByID", mock.Anything, uid).Return(mockProductResp, nil)

		ctx := context.TODO()
		p, err := mockProductService.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, p, mockProductResp)
		MockProductRepository.AssertExpectations(t)
	})
}
