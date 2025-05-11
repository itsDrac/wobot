package service

import (
	"github.com/itsDrac/wobot/store"
)

type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) UserService {
	return UserService{
		store: s,
	}
}
