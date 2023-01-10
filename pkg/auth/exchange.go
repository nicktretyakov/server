package auth

import (
	"context"

	"github.com/pascaldekloe/jwt"
	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/model"
	"be/pkg/xerror"
)

func (s *service) Exchange(ctx context.Context, code string) (*model.Session, *xerror.Error) {
	token, err := s.cfg.OauthCfg.Exchange(ctx, code)
	if err != nil {
		return nil, xerror.New(xerror.ProfileFailed, err.Error())
	}

	s.logger.Debug().Msgf("profile responded with access_token: %s", token.AccessToken)

	rawClaims, err := jwt.ParseWithoutCheck([]byte(token.AccessToken))
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	claims := claimsFromProfileToken(*rawClaims)
	if claims.employeePortalCode == 0 {
		return nil, xerror.New(xerror.PermissionDenied, "missing portal_code")
	}

	user, xerr := s.findOrCreateUser(ctx, claims)
	if xerr != nil {
		return nil, xerr
	}

	user.ProfileID = claims.profileID
	user.Phone = claims.phone
	user.Email = claims.email
	user.Employee = model.Employee{
		FirstName:  &claims.employeeFirstName,
		MiddleName: &claims.employeeMiddleName,
		LastName:   &claims.employeeLastName,
		Avatar:     &claims.employeeAvatar,
		Email:      &claims.employeeEmail,
		Phone:      &claims.employeePhone,
		Position:   &claims.employeePosition,
		PortalCode: claims.employeePortalCode,
	}

	user, xerr = s.updateUser(ctx, *user)
	if xerr != nil {
		return nil, xerr
	}

	return s.createSession(ctx, *user, *token)
}

func (s *service) findOrCreateUser(ctx context.Context, claims profileClaims) (*model.User, *xerror.Error) {
	u, err := s.users.FindUserByProfileID(ctx, claims.profileID)
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	if u != nil {
		return u, nil
	}

	u, err = s.users.FindUserByPortalCode(ctx, claims.employeePortalCode)
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	if u != nil {
		return u, nil
	}

	return s.createUser(ctx, claims)
}

func (s *service) createUser(ctx context.Context, claims profileClaims) (*model.User, *xerror.Error) {
	u, err := s.users.CreateUser(ctx, claims.asUser())
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	return u, nil
}

func (s *service) updateUser(ctx context.Context, u model.User) (*model.User, *xerror.Error) {
	user, err := s.users.UpdateUser(ctx, u)
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	return user, nil
}
