package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/muhammadderic/ecomrest/types"
)

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	handler := NewHandler(productStore)

	t.Run(
		"should succeed with valid product data",
		func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/products",
				nil,
			)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/api/v1/products", handler.handleGetProducts).Methods(http.MethodGet)
			router.ServeHTTP(rr, req)

			// Print the response recorder contents to the console
			t.Logf("Response Code: %d", rr.Code)
			t.Logf("Response Body: %s", rr.Body.String())

			if rr.Code != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
			}
		},
	)
}

type mockProductStore struct{}

func (m *mockProductStore) GetProducts() ([]*types.Product, error) {
	return nil, nil
}
