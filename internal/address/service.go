package address

import (
	"be/internal/datastore"
	"be/internal/user"
)

type Service struct {
	store       datastore.IDatastore
	userService user.IUserService
}

func New(store datastore.IDatastore, userService user.IUserService) *Service {
	s := &Service{
		store:       store,
		userService: userService,
	}

	return s
}
