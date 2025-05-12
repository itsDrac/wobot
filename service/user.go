package service

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/itsDrac/wobot/store"
	"github.com/itsDrac/wobot/types"
	"github.com/itsDrac/wobot/utils"
)

type UserService struct {
	store      store.Store
	jwtService JWTService
}

func NewUserService(s store.Store) UserService {
	jwtSecret := utils.GetStringEnv("JWT_SECRET", "secret")
	jwtExpiry := utils.GetIntEnv("JWT_EXPIRY", 3600)
	jwtExpiryTime := time.Now().Add(time.Duration(jwtExpiry) * time.Second)
	jwtService := NewJWTService(jwtSecret, jwtExpiryTime)
	return UserService{
		store:      s,
		jwtService: &jwtService,
	}
}

func GenreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s UserService) CreateUser(ctx context.Context, userPayload *types.CreateUserPayload) error {
	// Check if user exists with same username.
	var user store.User
	user.Username = userPayload.Username
	if err := s.store.GetUserByUsername(ctx, &user); err == nil {
		return fmt.Errorf("User with same Username exists")
	}
	// Hash the password.
	hash, err := GenreatePasswordHash(userPayload.Password)
	if err != nil {
		return fmt.Errorf("Can not hash password")
	}
	user.Password = hash
	// Create the user.
	if err := s.store.CreateUser(ctx, &user); err != nil {
		return fmt.Errorf("Can not create user %s", err.Error())
	}
	return nil
}

func (s UserService) LoginUser(ctx context.Context, userPayload *types.LoginUserPayload) (string, error) {
	// Get the user by username.
	var user store.User
	user.Username = userPayload.Username
	if err := s.store.GetUserByUsername(ctx, &user); err != nil {
		return "", fmt.Errorf("User with same Username does not exists")
	}
	// Check the password.
	if !CheckPasswordHash(userPayload.Password, user.Password) {
		return "", fmt.Errorf("Invalid password")
	}
	// Make JWT token.
	token, err := s.jwtService.GenerateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("Can not generate token")
	}
	// Return the token.
	return token, nil
}

func (s UserService) Authenticate(ctx context.Context, token string) (*store.User, error) {
	// Validate the token.
	username, err := s.jwtService.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}
	// Get the user by username.
	var user store.User
	user.Username = username
	if err := s.store.GetUserByUsername(ctx, &user); err != nil {
		return nil, fmt.Errorf("User with same Username does not exists")
	}
	return &user, nil
}
