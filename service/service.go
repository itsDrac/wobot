package service

import "github.com/itsDrac/wobot/store"

type Service struct {
	User interface {
	}
}

func NewService(s store.Store) Service {
	return Service{
		User: &UserService{
			store: s,
		},
	}
}
