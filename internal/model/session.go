package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UserID              uuid.UUID `json:"user_id" db:"user_id"`
	AccessToken         string    `json:"access_token" db:"access_token"`
	RefreshToken        string    `json:"refresh_token" db:"refresh_token"`
	RefreshExpires      time.Time `json:"refresh_expires" db:"refresh_expires"`
	ProfileAccessToken  string    `json:"profile_access_token" db:"profile_access_token"`
	ProfileRefreshToken string    `json:"profile_refresh_token" db:"profile_refresh_token"`
}
