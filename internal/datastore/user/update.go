package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/model"
)

func (s Storage) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	user.UpdatedAt = s.db.Now()

	sql := updateUserQuery(dbmodel.UserFromModel(user)).Where("id=?", user.ID)

	cmdTag, err := s.db.ExecBuilder(ctx, sql)
	if err != nil || cmdTag.RowsAffected() == 0 {
		if err != nil {
			return nil, err
		}

		return nil, base.ErrNotFound
	}

	return &user, nil
}

func updateUserQuery(u dbmodel.User) sq.UpdateBuilder {
	return base.Builder().
		Update(userTableName).
		SetMap(map[string]interface{}{
			"updated_at":           u.UpdatedAt,
			"portal_code":          u.PortalCode,
			"profile_id":           u.ProfileID,
			"email":                u.Email,
			"phone":                u.Phone,
			"employee_first_name":  u.EmployeeFirstName,
			"employee_middle_name": u.EmployeeMiddleName,
			"employee_last_name":   u.EmployeeLastName,
			"employee_avatar":      u.EmployeeAvatar,
			"employee_position":    u.EmployeePosition,
			"employee_email":       u.EmployeeEmail,
			"employee_phone":       u.EmployeePhone,
		})
}
