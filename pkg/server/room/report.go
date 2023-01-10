package room

import (
	"context"
	bookingpb "be/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) PrepareReportData(
	ctx context.Context,
	req *bookingpb.PrepareReportDataRequest,
) (*bookingpb.PrepareReportDataResponse, error) {
	roomID, err := uuid.Parse(req.GetRoomID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room id").Err()
	}

	timeline, err := getTimelineModel(req.GetTimeline())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	user := auth.FromContext(ctx)
	
	roo, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	if !roo.CanSendReport(*user) {
		return nil, status.New(codes.PermissionDenied, "permission denied").Err()
	}

	slots, equipments, releases, err := s.store.Room().FindRoomObjectsByTimeline(ctx, roomID, *timeline)
	if err != nil {
		return nil, err
	}

	return &bookingpb.PrepareReportDataResponse{
		Slots:      pbs.PbRoomSlotList(slots),
		Equipments: pbs.PbRoomEquipmentList(equipments),
		Releases:   pbs.PbRoomReleaseList(releases),
	}, nil
}
