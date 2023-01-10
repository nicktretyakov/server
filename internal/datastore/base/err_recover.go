package base

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const (
	UsersPortalCodeConstraint              = "users_portal_code_key"
	ReportsBookingPeriodKeyConstraint      = "booking_reports_booking_id_period_key"
	FinalReportsBookingPeriodKeyConstraint = "booking_final_reports_booking_id_key"
	BookingUsersBookingUserAssigneeUnique  = "booking_users_booking_user_role_2_idx"
)

var (
	ErrNotFound                    = errors.New("not found")
	ErrUserPortalCodeAlreadyExists = errors.New("user with portal code already exists")
	ErrBookingReportExists         = errors.New("booking report already exists")
	ErrBookingFinalReportExists    = errors.New("booking final report already exists")
	ErrBookingUserExists           = errors.New("booking user already exists")
)

func recoverUniqueConstraint(pgErr *pgconn.PgError) error {
	if pgErr.Code == pgerrcode.UniqueViolation {
		switch pgErr.ConstraintName {
		case UsersPortalCodeConstraint:
			return ErrUserPortalCodeAlreadyExists
		case ReportsBookingPeriodKeyConstraint:
			return ErrBookingReportExists
		case FinalReportsBookingPeriodKeyConstraint:
			return ErrBookingFinalReportExists
		case BookingUsersBookingUserAssigneeUnique:
			return ErrBookingUserExists
		}
	}

	return nil
}

func recoverDBError(err error) error {
	if err == nil {
		return nil
	}

	pgErr, ok := (err).(*pgconn.PgError)
	if ok {
		if tErr := recoverUniqueConstraint(pgErr); tErr != nil {
			return tErr
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return errors.Wrap(ErrNotFound, err.Error())
	}

	return err
}
