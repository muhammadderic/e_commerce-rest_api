package user

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	"github.com/muhammadderic/ecomrest/types"
// )

// func TestUserServiceHandlers(t *testing.T) {
// 	userStore := &mockUserStore{}
// 	handler := NewHandler(userStore)

// 	t.Run(
// 		"should succeed with valid user data",
// 		func(t *testing.T) {
// 			payload := types.RegisterUserPayload{
// 				FirstName: "John",
// 				LastName:  "Doe",
// 				Email:     "johndoe1@mail.com",
// 				Password:  "hispassword",
// 			}

// 			marshalled, _ := json.Marshal(payload)

// 			req, err := http.NewRequest(
// 				http.MethodPost,
// 				"/api/v1/register",
// 				bytes.NewBuffer(marshalled),
// 			)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()

// 			router := mux.NewRouter()
// 			router.HandleFunc("/api/v1/register", handler.handleRegister).Methods(http.MethodPost)

// 			router.ServeHTTP(rr, req)

// 			if rr.Code != http.StatusCreated {
// 				t.Errorf("expected %d, got %d", http.StatusCreated, rr.Code)
// 			} else {
// 				fmt.Println("Response Code:", rr.Code)
// 				fmt.Println("Response Body:", rr.Body.String())
// 			}
// 		},
// 	)
// }

// type mockUserStore struct{}

// func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
// 	// return nil, fmt.Errorf("user not found")
// 	if email == "johndoe1@mail.com" {
// 		return nil, fmt.Errorf("user not found")
// 	}
// 	return &types.User{Email: email}, nil
// }

// func (m *mockUserStore) CreateUser(types.User) error {
// 	return nil
// }

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/muhammadderic/ecomrest/types"
)

// Refactor the test to use the built-in testing package only
func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should succeed with valid user data", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe1@mail.com",
			Password:  "hispassword",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("failed to marshal payload: %v", err)
		}

		req, err := http.NewRequest(
			http.MethodPost,
			"/register",
			bytes.NewBuffer(marshalled),
		)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc(
			"/register",
			handler.handleRegister,
		)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	// Simulate that user does not exist
	if email == "johndoe1@mail.com" {
		return nil, nil
	}
	return &types.User{Email: email}, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
