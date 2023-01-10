package user

import (
	"context"

	"be/internal/datastore"
	"be/internal/model"
	"be/internal/profile"
)

type IUserService interface {
	GetOrCreateUserByPortalCode(ctx context.Context, portalCode uint64) (*model.User, error)
}

type Opts struct {
	Store       datastore.IUserStore
	EmployeeRep profile.IProfile
}

type service struct {
	store       datastore.IUserStore
	employeeRep profile.IProfile
}

func New(store datastore.IUserStore, employeeRep profile.IProfile) *service {
	return &service{store: store, employeeRep: employeeRep}
}
