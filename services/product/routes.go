package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadderic/ecomrest/types"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
}
