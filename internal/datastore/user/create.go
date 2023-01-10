package user

import (
	"context"

	"github.com/jackc/pgx/v4"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s Storage) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	if user.ID == uuid.Nil {
		user.ID = lib.UUID()
	}

	now := s.db.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Role = model.Regular

	ns := model.NoteSettings{
		UserID:  &user.ID,
		EmailOn: true,
		LifeOn:  true,
	}

	return &user, s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := s.db.ExecTxBuilder(ctx, tx, insertUserQuery(dbmodel.UserFromModel(user))); err != nil {
			return err
		}

		if _, err := s.db.ExecTxBuilder(ctx, tx, insertNoteSettingsQuery(ns)); err != nil {
			return err
		}

		return nil
	})
}

func insertUserQuery(u dbmodel.User) sq.InsertBuilder {
	return base.Builder().
		Insert(userTableName).
		Columns(
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
			"employee_email",
			"employee_phone",
		).
		Values(
			u.ID,         // id
			u.CreatedAt,  // created_at
			u.UpdatedAt,  // updated_at
			u.PortalCode, // portal_code
			u.ProfileID,  // profile_id
			u.Email,      // email
			u.Phone,      // phone
			u.Role,       // role
			u.EmployeeFirstName,
			u.EmployeeMiddleName,
			u.EmployeeLastName,
			u.EmployeeAvatar,
			u.EmployeeEmail,
			u.EmployeePhone,
		)
}
