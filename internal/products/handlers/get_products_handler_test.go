package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"products/internal/products/handlers/mocks"
	"products/internal/products/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Should retrieve a product", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockProduct := &model.Product{
			UID: uid,
		}

		mockProductService := new(mocks.MockProductService)
		mockProductService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockProduct, nil)

		responseRecorder := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			ctx.Set("product", &model.Product{
				UID: uid,
			})
		})

		NewHandler(&Config{
			R:              router,
			ProductService: mockProductService,
		})
		url := "/api/v1/products/" + uid.String()
		request, err := http.NewRequest(http.MethodGet, url, nil)
		assert.NoError(t, err)

		router.ServeHTTP(responseRecorder, request)
		respBody, err := json.Marshal(gin.H{
			"product": mockProduct,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, responseRecorder.Code)
		assert.Equal(t, respBody, responseRecorder.Body.Bytes())
		mockProductService.AssertExpectations(t)
	})
}
