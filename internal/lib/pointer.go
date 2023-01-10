package lib

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

func String(v string) *string {
	return &v
}

func PUUID(v uuid.UUID) *uuid.UUID {
	return &v
}

func Uint(v uint) *uint {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func Notification(v model.Notification) *model.Notification {
	return &v
}
