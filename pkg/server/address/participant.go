package address

import (
	"context"
	"unicode/utf8"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/pkg/auth"
	"be/pkg/server/pbs"
)

const (
	minRoleLen = 3
	limitRole  = 70
)

func (s Service) AddParticipant(ctx context.Context, req *bookingpb.AddParticipantRequest) (*bookingpb.AddParticipantResponse, error) {
	addressID, role, err := validateRequest(req.GetAddressID(), req.GetRole())
	if err != nil {
		return nil, err
	}

	participant, role, err := s.addressService.AddParticipant(
		ctx,
		*auth.FromContext(ctx),
		role,
		uint64(req.GetPortalCodeParticipant()),
		*addressID,
		getTypeAddress(req.GetAddressType()),
	)
	if err != nil {
		return nil, err
	}

	return &bookingpb.AddParticipantResponse{
		Participant: pbs.PbParticipant(participant, role),
	}, nil
}

func (s Service) UpdateParticipant(
	ctx context.Context,
	req *bookingpb.UpdateParticipantRequest,
) (*bookingpb.UpdateParticipantResponse, error) {
	addressID, role, err := validateRequest(req.GetAddressID(), req.GetRole())
	if err != nil {
		return nil, err
	}

	participant, role, err := s.addressService.UpdateParticipant(
		ctx,
		*auth.FromContext(ctx),
		role,
		uint64(req.GetPortalCodeParticipantBase()),
		uint64(req.GetPortalCodeParticipantChange()),
		*addressID,
		getTypeAddress(req.GetAddressType()),
	)
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateParticipantResponse{
		Participant: pbs.PbParticipant(participant, role),
	}, nil
}

func (s Service) RemoveParticipant(ctx context.Context, req *bookingpb.RemoveParticipantRequest) (*emptypb.Empty, error) {
	addressID, err := uuid.Parse(req.GetAddressId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid invest object id")
	}

	if err = s.addressService.RemoveParticipant(
		ctx,
		*auth.FromContext(ctx),
		uint64(req.GetPortalCodeParticipant()),
		addressID,
		getTypeAddress(req.GetAddressType()),
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func validateRequest(addressUUID, role string) (*uuid.UUID, string, error) {
	addressID, err := uuid.Parse(addressUUID)
	if err != nil {
		return nil, "", status.Error(codes.InvalidArgument, "invalid invest object id")
	}

	if utf8.RuneCountInString(role) < minRoleLen || utf8.RuneCountInString(role) > limitRole {
		return nil, "", status.Error(codes.InvalidArgument, "invalid role")
	}

	return &addressID, role, nil
}
