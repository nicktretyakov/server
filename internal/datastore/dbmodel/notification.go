package dbmodel

import (
	"github.com/jackc/pgtype"

	"be/internal/model"
)

func ToNotification(p pgtype.Numeric) model.Notification {
	if p.Status == pgtype.Undefined {
		return model.NewNotificationFromFloat(0)
	}

	return model.NewNotification(p.Int.Int64(), p.Exp)
}

func ToNotificationPtr(p pgtype.Numeric) *model.Notification {
	if p.Status == pgtype.Null {
		return nil
	}

	if p.Status == pgtype.Undefined {
		return nil
	}

	m := model.NewNotification(p.Int.Int64(), p.Exp)

	return &m
}
