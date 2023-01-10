package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pascaldekloe/jwt"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"

	"be/internal/lib"
	"be/internal/model"
	"be/pkg/xerror"
)

const refreshLength = 100

type service struct {
	cfg    Config
	users  IAuthUserStore
	logger zerolog.Logger
}

type Config struct {
	OauthCfg      oauth2.Config
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
	Secret        []byte
}

func New(cfg Config, store IAuthUserStore, logger zerolog.Logger) IAuth {
	return &service{
		cfg:    cfg,
		users:  store,
		logger: logger,
	}
}

func (s *service) AuthURL() string {
	return s.cfg.OauthCfg.AuthCodeURL("")
}

func (s *service) ParseToken(ctx context.Context, token string) (*model.User, *model.Session, *xerror.Error) {
	clm, err := jwt.HMACCheck([]byte(token), s.cfg.Secret)
	if err != nil {
		return nil, nil, xerror.Newf(xerror.UnAuthorized, "failed to parse token: %s", err.Error())
	}

	if !clm.Valid(time.Now()) {
		return nil, nil, xerror.New(xerror.UnAuthorized, "expired token")
	}

	sessionID, err := uuid.Parse(clm.Subject)
	if err != nil {
		return nil, nil, xerror.Newf(xerror.UnAuthorized, "failed to parse token: %s", err.Error())
	}

	session, err := s.users.FindSessionByID(ctx, sessionID)
	if err != nil {
		return nil, nil, xerror.New(xerror.Internal, err.Error())
	}

	user, err := s.users.FindUserByPK(ctx, session.UserID)
	if err != nil {
		return nil, nil, xerror.New(xerror.Internal, err.Error())
	}

	return user, session, nil
}

func (s *service) createSession(ctx context.Context,
	user model.User, profileToken oauth2.Token,
) (*model.Session, *xerror.Error) {
	now := time.Now()
	refreshExpires := now.Add(s.cfg.RefreshExpiry)

	sessionID := lib.UUID()

	accessToken, err := s.accessToken(sessionID, user)
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	refreshToken, err := s.refreshToken()
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	session := model.Session{
		ID:                  sessionID,
		CreatedAt:           now,
		UserID:              user.ID,
		AccessToken:         *accessToken,
		RefreshToken:        refreshToken,
		RefreshExpires:      refreshExpires,
		ProfileAccessToken:  profileToken.AccessToken,
		ProfileRefreshToken: profileToken.RefreshToken,
	}

	createdSession, err := s.users.NewSession(ctx, session)
	if err != nil {
		return nil, xerror.New(xerror.Internal, err.Error())
	}

	return createdSession, nil
}

func (s *service) refreshToken() (string, error) {
	return lib.GenerateRandomString(refreshLength)
}

func (s *service) accessToken(sessionID uuid.UUID, user model.User) (*string, error) {
	now := time.Now()
	accessExpires := now.Add(s.cfg.AccessExpiry)

	claims := jwt.Claims{
		Registered: jwt.Registered{
			Issued:  jwt.NewNumericTime(now.Truncate(time.Second)),
			Expires: jwt.NewNumericTime(accessExpires.Truncate(time.Second)),
			Subject: sessionID.String(),
		},
		Set: map[string]interface{}{
			"user_id":   user.ID.String(),
			"user_role": user.Role,
		},
	}

	token, err := claims.HMACSign(jwt.HS512, s.cfg.Secret)
	if err != nil {
		return nil, err
	}

	tokenStr := string(token)

	return &tokenStr, nil
}
