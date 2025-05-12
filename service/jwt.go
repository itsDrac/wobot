package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}

type JWTServiceImpl struct {
	secretKey string
	expiry    int64
}

func NewJWTService(secretKey string, expiry time.Time) JWTServiceImpl {
	return JWTServiceImpl{
		secretKey: secretKey,
		expiry:    expiry.Unix(),
	}
}

func (j *JWTServiceImpl) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      j.expiry,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JWTServiceImpl) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}
	return username, nil
}
