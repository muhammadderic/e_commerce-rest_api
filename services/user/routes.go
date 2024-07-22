package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/muhammadderic/ecomrest/services/auth"
	"github.com/muhammadderic/ecomrest/types"
	"github.com/muhammadderic/ecomrest/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid request: %s", errors),
		)
		return
	}

	// Check if the user exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("user %s already exists", user.Email),
		)
		return
	}

	// Hash the user's password
	hashedPasword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	// Create the user
	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPasword,
	})
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)
		return
	}

	// Return a HTTP 201 status code if the user is created successfully
	utils.WriteJSON(w, http.StatusCreated, nil)
}
