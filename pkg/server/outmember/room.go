package outmember

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	addressStore "be/internal/datastore/address"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) AcceptRoomAsRegister(ctx context.Context, req *bookingpb.RoomRequest) (*bookingpb.RoomResponse, error) {
	return s.newInitialReportForRoom(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineRoomAsRegister(ctx context.Context, req *bookingpb.RoomRequest) (*bookingpb.RoomResponse, error) {
	return s.newInitialReportForRoom(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) AcceptRoomAsAssignee(ctx context.Context, req *bookingpb.RoomRequest) (*bookingpb.RoomResponse, error) {
	return s.newApprovalReportForRoom(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineRoomAsAssignee(ctx context.Context, req *bookingpb.RoomRequest) (*bookingpb.RoomResponse, error) {
	return s.newApprovalReportForRoom(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) newInitialReportForRoom(ctx context.Context, req *bookingpb.RoomRequest,
	result model.OutmemberResult,
) (*bookingpb.RoomResponse, error) {
	roomID, err := uuid.Parse(req.GetRoomID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room_id").Err()
	}

	user := auth.FromContext(ctx)

	if !acl.IsHeadOfBooking(*user) {
		return nil, status.New(codes.PermissionDenied, "superuser only").Err()
	}

	if err = s.outmemberService.NewInitialOutmember(ctx, *user, addressStore.RoomAddressType, model.Outmember{
		Type:           model.InitialOutmemberType,
		UserID:         user.ID,
		AddressID: roomID,
		Result:         result,
		Extra:          map[string]interface{}{"comment": req.GetComment()},
		Role:           model.BookingManagerOutmemberRole,
	}); err != nil {
		return nil, err
	}

	room, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateRoomNote(user,
		getNoteTypeByStatus(room.Status),
		room,
	)

	return &bookingpb.RoomResponse{Room: pbs.PbRoom(*room)}, nil
}

func (s Service) newApprovalReportForRoom(ctx context.Context, req *bookingpb.RoomRequest,
	result model.OutmemberResult,
) (*bookingpb.RoomResponse, error) {
	roomID, err := uuid.Parse(req.GetRoomID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room_id").Err()
	}

	user := auth.FromContext(ctx)

	if err = s.outmemberService.NewAcceptanceOutmember(ctx, *user, addressStore.RoomAddressType, model.Outmember{
		Type:           model.ApprovalOutmemberType,
		UserID:         user.ID,
		AddressID: roomID,
		Result:         result,
		Extra:          map[string]interface{}{"comment": req.GetComment()},
		Role:           model.AssigneeOutmemberRole,
	}); err != nil {
		return nil, err
	}

	room, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateRoomNote(user,
		getNoteTypeByStatus(room.Status),
		room,
	)

	return &bookingpb.RoomResponse{Room: pbs.PbRoom(*room)}, nil
}
