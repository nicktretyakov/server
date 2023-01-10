package auth

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
	"be/pkg/xerror"
)

type Token struct {
	AccessToken string
}

type IAuth interface {
	AuthURL() string
	Exchange(ctx context.Context, code string) (*model.Session, *xerror.Error)
	Refresh(ctx context.Context, token string) (*model.Session, *xerror.Error)
	ParseToken(ctx context.Context, token string) (*model.User, *model.Session, *xerror.Error)
}

type IAuthUserStore interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	FindUserByProfileID(ctx context.Context, uid string) (*model.User, error)
	FindUserByPortalCode(ctx context.Context, uid uint64) (*model.User, error)
	FindUserByPK(ctx context.Context, userID uuid.UUID) (*model.User, error)
	EmployeeInfo(ctx context.Context, portalCode uint64) (*model.Employee, error)
	NewSession(ctx context.Context, session model.Session) (*model.Session, error)
	FindSessionByRefreshToken(ctx context.Context, refresh string) (*model.Session, error)
	FindSessionByID(ctx context.Context, sessionID uuid.UUID) (*model.Session, error)
	RefreshSession(ctx context.Context, session model.Session) error
}
