package booking

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s *Storage) CreateBookingUser(ctx context.Context, bookingUser model.BookingUser) error {
	_, err := s.db.ExecBuilder(ctx, bookingUserInsertQuery(bookingUser))

	return err
}

func (s *Storage) UpdateBookingUser(ctx context.Context, bookingUserBase, bookingUserChange model.BookingUser) error {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if err := s.DeleteBookingUser(ctx, bookingUserBase); err != nil {
			return err
		}

		if err := s.CreateBookingUser(ctx, bookingUserChange); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteBookingUser(ctx context.Context, bookingUser model.BookingUser) error {
	cmd, err := s.db.ExecBuilder(ctx, bookingUserDeleteQuery(bookingUser))
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return base.ErrNotFound
	}

	return err
}

func (s *Storage) GetListSupervisors(ctx context.Context) ([]uuid.UUID, error) {
	type bookingUserTable struct {
		BookingID uuid.UUID             `db:"booking_id"`
		Role      model.BookingUserRole `db:"role"`
		UserID    uuid.UUID             `db:"user_id"`
	}

	var bookingUsers []bookingUserTable

	if err := s.db.Select(ctx, bookingUserSupervisorQuery(), &bookingUsers); err != nil {
		return nil, err
	}

	userIDs := make([]uuid.UUID, 0)

	for _, user := range bookingUsers {
		userIDs = append(userIDs, user.UserID)
	}

	return userIDs, nil
}

func bookingUserInsertQuery(bookingUser model.BookingUser) sq.InsertBuilder {
	return base.Builder().Insert(bookingUserRoleTableName).
		Columns(
			"booking_id",
			"user_id",
			"role",
			"role_description",
		).
		Values(
			bookingUser.BookingID,
			bookingUser.UserID,
			bookingUser.Role,
			bookingUser.RoleDescription,
		)
}

func bookingUserDeleteQuery(bookingUser model.BookingUser) sq.DeleteBuilder {
	return base.Builder().Delete(bookingUserRoleTableName).
		Where("booking_id=? and user_id = ? and role = ?", bookingUser.BookingID, bookingUser.UserID, bookingUser.Role)
}

func bookingUserSupervisorQuery() sq.SelectBuilder {
	return base.Builder().Select("user_id").From(bookingUserRoleTableName).GroupBy("user_id").
		Where("role = ?", model.SupervisorBookingUserRole)
}
