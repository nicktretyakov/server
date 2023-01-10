package room

import (
	"be/internal/datastore"
	"be/internal/user"
)

type Service struct {
	store       datastore.IDatastore
	userService user.IUserService
}

func New(store datastore.IDatastore, userService user.IUserService) *Service {
	return &Service{
		store:       store,
		userService: userService,
	}
}
