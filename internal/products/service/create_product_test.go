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

func TestCreate(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		//generamos un UUID random
		uid, _ := uuid.NewRandom()
		//mockeamos la respuesta del FindByID
		mockProductResp := &model.Product{
			UID: uid,
		}
		//mockeamos el repository
		mockProductRepository := new(mocks.MockProductRepository)
		mockProductService := NewProductService(&ProductConfig{
			ProductRepository: mockProductRepository,
		})
		mockProductRepository.
			On("CreateProduct", mock.AnythingOfType("*context.emptyCtx"), mockProductResp).
			Run(func(args mock.Arguments) {
				userArg := args.Get(1).(*model.Product) // arg 0 is context, arg 1 is *User
				userArg.UID = uid
			}).Return(nil)
		ctx := context.TODO()
		err := mockProductService.Create(ctx, mockProductResp)

		assert.NoError(t, err)
		assert.Equal(t, uid, mockProductResp.UID)
		mockProductRepository.AssertExpectations(t)
	})
}
