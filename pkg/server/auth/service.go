package auth

import (
	bookingpb "be/proto"

	"be/pkg/auth"
)

type Service struct {
	bookingpb.UnimplementedAuthServer
	authService auth.IAuth
}

func NewService(authService auth.IAuth) bookingpb.AuthServer {
	return &Service{authService: authService}
}
