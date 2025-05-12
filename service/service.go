package service

import (
	"context"
	"mime/multipart"

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
		UploadFile(ctx context.Context, file multipart.File, fileInfo *multipart.FileHeader) error
		GetRemainingStorage(ctx context.Context) (string, error)
		GetFiles(ctx context.Context) (types.Files, error)
	}
}

func NewService(s store.Store) Service {
	UserService := NewUserService(s)
	FileService := NewFileService(s)
	return Service{
		User: &UserService,
		File: &FileService,
	}
}
