package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/itsDrac/wobot/types"
)

var (
	ErrInvalidRequestData = errors.New("invalid request data")
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user by providing username and password
// @Tags Users
// @Accept json
// @Param data body types.CreateUserPayload true "User data"
// @Success 201 {object} types.ApiResponse
// @Failure 400 {object} string "Invalid request data"
// @Failure 409 {object} string
// @Failure 500 {object} string "Internal server error"
// @Router /users/create [post]
func (h ChiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user types.CreateUserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, ErrInvalidRequestData.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request body
	if err := validate.Struct(user); err != nil {
		http.Error(w, ErrInvalidRequestData.Error(), http.StatusBadRequest)
		return
	}

	// Create the user
	if err := h.Service.User.CreateUser(r.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "User with same Username exists") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if strings.Contains(err.Error(), "Can not hash password") {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		if strings.Contains(err.Error(), "Can not create user") {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error creating user:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := types.ApiResponse{
		Message: "User created successfully",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}

// LoginUser godoc
// @Summary Login an existing user
// @Description Login an existing user by providing username and password
// @Tags Users
// @Accept json
// @Param data body types.LoginUserPayload true "User data"
// @Success 200 {object} types.LoginUserResponsePayload
// @Failure 400 {object} string "Invalid request data"
// @Failure 404 {object} string "User not found"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /users/login [post]
func (h ChiHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user types.LoginUserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, ErrInvalidRequestData.Error(), http.StatusBadRequest)
		return
	}
	// Validate the request body
	if err := validate.Struct(user); err != nil {
		http.Error(w, ErrInvalidRequestData.Error(), http.StatusBadRequest)
		return
	}

	// Login the user
	token, err := h.Service.User.LoginUser(r.Context(), &user)
	if err != nil {
		if strings.Contains(err.Error(), "User with same Username exists") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if strings.Contains(err.Error(), "Invalid password") {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		}
		if strings.Contains(err.Error(), "Can not generate token") {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error login user:", err)
		return
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	resp := types.LoginUserResponsePayload{
		Token: token,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}
}
