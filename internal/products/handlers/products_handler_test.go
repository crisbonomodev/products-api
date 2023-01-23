package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetProducts(t *testing.T) {
// 	t.Run("Should return a product", func(t *testing.T) {
// 		server := InitializeServer()

// 		r, _ := http.NewRequest(http.MethodGet, "/api/v1/products/12345678", nil)
// 		w := httptest.NewRecorder()

// 		server.ServeHTTP(w, r)

// 		var response map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &response)

// 		assert.Nil(t, err)
// 		assert.Equal(t, 200, w.Code)
// 		assert.Equal(t, "12345678", response["id"])
// 	})
// }

// func TestCreateProducts(t *testing.T) {
// 	t.Run("Should parse a body and create a product", func(t *testing.T) {
// 		r, _ := http.NewRequest(http.MethodPost, "/api/v1/products", nil)
// 		server := InitializeServer()
// 		w := httptest.NewRecorder()

// 		server.ServeHTTP(w, r)

// 		var response map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &response)

// 		assert.Nil(t, err)
// 		assert.Equal(t, 201, w.Code)

// 	})
// }
