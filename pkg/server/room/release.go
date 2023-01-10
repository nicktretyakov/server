package room

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
	"be/pkg/server/validators"
)

func (s Service) AddRelease(ctx context.Context, req *bookingpb.AddReleaseRequest) (*bookingpb.AddReleaseResponse, error) {
	releasesFromPB, err := getReleases(req.GetReleases())
	if err != nil {
		return nil, err
	}

	if releasesFromPB == nil {
		return &bookingpb.AddReleaseResponse{}, nil
	}

	room, err := s.store.Room().FindByID(ctx, releasesFromPB[0].RoomID)
	if err != nil {
		return nil, err
	}

	if !room.CanUpdateRelease(*auth.FromContext(ctx)) {
		return nil, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	releases, err := s.store.Room().AddReleases(ctx, releasesFromPB)
	if err != nil {
		return nil, err
	}

	releasesIDs := make([]string, 0, len(releases))

	for _, release := range releases {
		releasesIDs = append(releasesIDs, release.ID.String())
	}

	return &bookingpb.AddReleaseResponse{Uuids: releasesIDs}, nil
}

func (s Service) UpdateRelease(ctx context.Context, req *bookingpb.UpdateReleaseRequest) (*bookingpb.UpdateReleaseResponse, error) {
	forUpdateRelease, err := validateRelease(req.GetRelease())
	if err != nil {
		return nil, err
	}

	room, err := s.store.Room().FindByID(ctx, forUpdateRelease.RoomID)
	if err != nil {
		return nil, err
	}

	if !room.CanUpdateRelease(*auth.FromContext(ctx)) {
		return nil, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	release, err := s.store.Room().UpdateRelease(ctx, *forUpdateRelease)
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateReleaseResponse{
		Release: pbs.PbRoomRelease(*release),
	}, nil
}

func (s Service) RemoveRelease(ctx context.Context, req *bookingpb.RemoveReleaseRequest) (*emptypb.Empty, error) {
	releaseID, err := uuid.Parse(req.GetUuid())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid release id").Err()
	}

	release, err := s.store.Room().FindReleaseByID(ctx, releaseID)
	if err != nil {
		return nil, err
	}

	room, err := s.store.Room().FindByID(ctx, release.RoomID)
	if err != nil {
		return nil, err
	}

	if !room.CanUpdateRelease(*auth.FromContext(ctx)) {
		return nil, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	if err = s.store.Room().DeleteRelease(ctx, releaseID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func getReleases(releases []*bookingpb.Release) ([]model.Release, error) {
	if len(releases) == 0 {
		return nil, nil
	}

	modelReleases := make([]model.Release, 0, len(releases))

	for _, pbRelease := range releases {
		validRelease, err := validateRelease(pbRelease)
		if err != nil {
			return nil, err
		}

		modelReleases = append(modelReleases, *validRelease)
	}

	return modelReleases, nil
}

func validateRelease(pbRelease *bookingpb.Release) (*model.Release, error) {
	var (
		releaseUUID uuid.UUID
		err         error
	)

	if pbRelease.GetUuid() != "" {
		releaseUUID, err = uuid.Parse(pbRelease.GetUuid())
		if err != nil {
			return nil, status.New(codes.InvalidArgument, "invalid release id").Err()
		}
	}

	roomID, err := uuid.Parse(pbRelease.GetRoomID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room id").Err()
	}

	if len(pbRelease.GetTitle()) < minTitleLen {
		return nil, status.New(codes.InvalidArgument, "invalid title").Err()
	}

	if len(pbRelease.GetDescription()) < minDescriptionLen {
		return nil, status.New(codes.InvalidArgument, "invalid description").Err()
	}

	date, err := validators.Time(pbRelease.GetDate())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid date").Err()
	}

	factSlot, err := validators.Notification(pbRelease.GetFactSlot())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid fact slot").Err()
	}

	return &model.Release{
		ID:          releaseUUID,
		RoomID:   roomID,
		Title:       pbRelease.GetTitle(),
		Description: pbRelease.GetDescription(),
		Date:        date,
		FactSlot:  factSlot,
	}, nil
}
