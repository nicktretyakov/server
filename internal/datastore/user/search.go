package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s Storage) FindAdmins(ctx context.Context) ([]*model.User, error) {
	var users dbmodel.UserList

	err := s.db.Select(ctx, SelectUserQuery().Where(sq.Eq{"role": model.Admin}).Where(sq.NotEq{"email": ""}), &users)
	if err != nil {
		return nil, err
	}

	return users.UsersPtr(), nil
}

func (s Storage) FindUserByProfileID(ctx context.Context, uid string) (*model.User, error) {
	return findOne(ctx, s.db, "profile_id=?", uid)
}

func (s Storage) FindUserByPortalCode(ctx context.Context, portalCode uint64) (*model.User, error) {
	return findOne(ctx, s.db, "portal_code=?", portalCode)
}

func (s Storage) FindUserByPK(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return findOne(ctx, s.db, "id=?", userID)
}

func (s Storage) FindUsersByIDs(ctx context.Context, userIDs []uuid.UUID) ([]model.User, error) {
	return findIn(ctx, s.db, sq.Eq{"id": userIDs})
}

func (s Storage) FindUsersIDByPortalCode(ctx context.Context, uid []uint64) ([]uuid.UUID, error) {
	var usersID []uuid.UUID

	err := s.db.Select(ctx, selectUsersIDQuery(uid), &usersID)
	if err != nil {
		return nil, err
	}

	return usersID, nil
}

func findOne(ctx context.Context, db *base.DB, pred interface{}, args ...interface{}) (*model.User, error) {
	var user dbmodel.User

	err := db.Get(ctx, SelectUserQuery().Where(pred, args...), &user)
	if err != nil {
		return nil, err
	}

	return user.ToModelPtr(), nil
}

func findIn(ctx context.Context, db *base.DB, pred interface{}) ([]model.User, error) {
	var users dbmodel.UserList

	err := db.Select(ctx, SelectUserQuery().Where(pred), &users)
	if err != nil {
		return nil, err
	}

	return users.Users(), nil
}

func SelectUserQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"created_at",
			"updated_at",
			"portal_code",
			"profile_id",
			"email",
			"phone",
			"role",
			"employee_first_name",
			"employee_middle_name",
			"employee_last_name",
			"employee_avatar",
			"employee_position",
			"employee_email",
			"employee_phone",
		).
		From(userTableName)
}

func selectUsersIDQuery(uid []uint64) sq.SelectBuilder {
	return base.Builder().Select("id").From(userTableName).Where(sq.Eq{"portal_code": uid})
}
