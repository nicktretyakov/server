package auth

import (
	"context"
	"time"

	"be/internal/model"
	"be/pkg/xerror"
)

func (s *service) Refresh(ctx context.Context, refreshToken string) (*model.Session, *xerror.Error) {
	session, err := s.users.FindSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, xerror.Newf(xerror.Internal, "session by token: %s", err.Error())
	}

	user, err := s.users.FindUserByPK(ctx, session.UserID)
	if err != nil {
		return nil, xerror.Newf(xerror.Internal, "user by pk: %s", err.Error())
	}

	newAccessToken, err := s.accessToken(session.ID, *user)
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	newRefreshToken, err := s.refreshToken()
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	session.AccessToken = *newAccessToken
	session.RefreshToken = newRefreshToken
	session.RefreshExpires = time.Now().Add(s.cfg.RefreshExpiry)

	if err = s.users.RefreshSession(ctx, *session); err != nil {
		return nil, xerror.Newf(xerror.Internal, "refresh session: %s", err.Error())
	}

	return session, nil
}
