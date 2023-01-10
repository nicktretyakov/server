package auth

import (
	"context"

	bookingpb "be/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Service) AuthByCode(ctx context.Context,
	request *bookingpb.AuthByCodeRequest,
) (*bookingpb.AuthByCodeResponse, error) {
	session, aErr := s.authService.Exchange(ctx, request.GetCode())
	if aErr != nil {
		return nil, aErr
	}

	return &bookingpb.AuthByCodeResponse{
		Access:  session.AccessToken,
		Refresh: session.RefreshToken,
	}, nil
}

func (s Service) RefreshToken(ctx context.Context,
	request *bookingpb.RefreshTokenRequest,
) (*bookingpb.RefreshTokenResponse, error) {
	session, aErr := s.authService.Refresh(ctx, request.GetRefresh())
	if aErr != nil {
		return nil, aErr
	}

	return &bookingpb.RefreshTokenResponse{
		Access:  session.AccessToken,
		Refresh: session.RefreshToken,
	}, nil
}

func (s Service) AuthURL(context.Context, *emptypb.Empty) (*bookingpb.AuthURLResponse, error) {
	return &bookingpb.AuthURLResponse{
		AuthUrl: s.authService.AuthURL(),
	}, nil
}
