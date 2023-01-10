package commands

import (
	"context"

	"github.com/google/uuid"

	"be/internal/datastore"
	"be/internal/model"
	"be/internal/profile"
	"be/pkg/auth"
)

// Think about it!
type adapter struct {
	store      datastore.IDatastore
	profileAPI profile.IProfile
}

func (a adapter) EmployeeInfo(ctx context.Context, portalCode uint64) (*model.Employee, error) {
	resp, err := a.profileAPI.FindEmployeeByPortalCode(ctx, portalCode)
	if err != nil {
		return nil, err
	}

	emp := resp.Cast()

	return &emp, nil
}

func (a adapter) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	return a.store.User().CreateUser(ctx, user)
}

func (a adapter) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	return a.store.User().UpdateUser(ctx, user)
}

func (a adapter) FindUserByProfileID(ctx context.Context, uid string) (*model.User, error) {
	return a.store.User().FindUserByProfileID(ctx, uid)
}

func (a adapter) FindUserByPortalCode(ctx context.Context, uid uint64) (*model.User, error) {
	return a.store.User().FindUserByPortalCode(ctx, uid)
}

func (a adapter) FindUserByPK(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return a.store.User().FindUserByPK(ctx, userID)
}

func (a adapter) FindUsersByIDs(ctx context.Context, userIDs []uuid.UUID) ([]model.User, error) {
	return a.store.User().FindUsersByIDs(ctx, userIDs)
}

func (a adapter) NewSession(ctx context.Context, session model.Session) (*model.Session, error) {
	return a.store.Session().NewSession(ctx, session)
}

func (a adapter) FindSessionByRefreshToken(ctx context.Context, refresh string) (*model.Session, error) {
	return a.store.Session().FindSessionByRefreshToken(ctx, refresh)
}

func (a adapter) FindSessionByID(ctx context.Context, sessionID uuid.UUID) (*model.Session, error) {
	return a.store.Session().FindSessionByID(ctx, sessionID)
}

func (a adapter) RefreshSession(ctx context.Context, session model.Session) error {
	return a.store.Session().RefreshSession(ctx, session)
}

func toAuthUser(store datastore.IDatastore, profileAPI profile.IProfile) auth.IAuthUserStore {
	return &adapter{store: store, profileAPI: profileAPI}
}
