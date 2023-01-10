package user

import (
	"context"

	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s service) GetOrCreateUserByPortalCode(ctx context.Context, portalCode uint64) (*model.User, error) {
	user, err := s.store.FindUserByPortalCode(ctx, portalCode)
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	employee, err := s.employeeRep.FindEmployeeByPortalCode(ctx, portalCode)
	if err != nil {
		return nil, err
	}

	userToCreate := model.User{
		ProfileID: "",
		Email:     "",
		Phone:     "",
		Role:      model.Regular,
		Employee:  employee.Cast(),
	}

	userToCreate.Employee.PortalCode = portalCode

	if employee.ProfileID != nil {
		userToCreate.ProfileID = employee.ProfileID.String()
	}

	return s.store.CreateUser(ctx, userToCreate)
}
