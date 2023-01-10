package address

import (
	"context"
	"regexp"
	"unicode/utf8"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

const minLinkNameLen = 3

func (s Service) AddLink(ctx context.Context, req *bookingpb.AddLinkRequest) (*bookingpb.AddLinkResponse, error) {
	addressID, err := uuid.Parse(req.GetAddressID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address_id")
	}

	user := auth.FromContext(ctx)

	if err := validateLinkRequest(req.GetName(), req.GetSource()); err != nil {
		return nil, err
	}

	link := model.Link{
		Name:   req.GetName(),
		Source: req.GetSource(),
	}

	if err = s.addressService.AddLink(ctx, *user, &link, addressID, getTypeAddress(req.GetAddressType())); err != nil {
		return nil, err
	}

	return &bookingpb.AddLinkResponse{
		Link: pbs.PbLink(link),
	}, nil
}

func (s Service) UpdateLink(ctx context.Context, req *bookingpb.UpdateLinkRequest) (*bookingpb.UpdateLinkResponse, error) {
	addressID, err := uuid.Parse(req.GetAddressID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address_id")
	}

	user := auth.FromContext(ctx)

	if req.GetLink() == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid link")
	}

	linkUUID, err := uuid.Parse(req.GetLink().GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid link_id")
	}

	if err := validateLinkRequest(req.GetLink().GetName(), req.GetLink().GetSource()); err != nil {
		return nil, err
	}

	link := model.Link{
		Id:     linkUUID,
		Name:   req.GetLink().GetName(),
		Source: req.GetLink().GetSource(),
	}

	if err = s.addressService.UpdateLink(ctx, *user, &link, addressID, getTypeAddress(req.GetAddressType())); err != nil {
		return nil, err
	}

	return &bookingpb.UpdateLinkResponse{
		Link: pbs.PbLink(link),
	}, nil
}

func (s Service) RemoveLink(ctx context.Context, req *bookingpb.RemoveLinkRequest) (*emptypb.Empty, error) {
	addressID, err := uuid.Parse(req.GetAddressID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address_id")
	}

	user := auth.FromContext(ctx)

	linkUUID, err := uuid.Parse(req.GetLinkId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid link_id")
	}

	if err = s.addressService.
		RemoveLink(
			ctx,
			*user,
			linkUUID,
			addressID,
			getTypeAddress(req.GetAddressType()),
		); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func validateLinkRequest(name string, source string) error {
	if utf8.RuneCountInString(name) < minLinkNameLen {
		return status.New(codes.InvalidArgument, "invalid name").Err()
	}

	match, err := regexp.MatchString(`^(https?://).+\..{2,}`, source)
	if err != nil {
		return status.New(codes.Internal, "error matching regexp").Err()
	}

	if !match {
		return status.New(codes.InvalidArgument, "invalid url format").Err()
	}

	return nil
}
