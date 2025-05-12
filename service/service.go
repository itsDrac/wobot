package service

import (
	"context"

	"github.com/itsDrac/wobot/store"
	"github.com/itsDrac/wobot/types"
)

type Service struct {
	User interface {
		CreateUser(ctx context.Context, userPayload *types.CreateUserPayload) error
		LoginUser(ctx context.Context, userPayload *types.LoginUserPayload) (string, error)
		Authenticate(ctx context.Context, token string) (*store.User, error)
	}
	File interface {
		UploadFile(ctx context.Context, filePayload *types.UploadFilePayload) error
	}
}

func NewService(s store.Store) Service {
	UserService := NewUserService(s)
	return Service{
		User: &UserService,
	}
}
