package room

import (
	"context"
	"time"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/model"
	"be/pkg/auth"
)

func (s Service) UpdateRoom(ctx context.Context, req *bookingpb.UpdateRoomRequest) (*bookingpb.UpdateRoomResponse, error) { //nolint:lll
	updReq, equipmentIDs, slotIDs, bookingIDs, err := validateUpdateRequest(req)
	if err != nil {
		return nil, err
	}

	roomToUpdate, err := s.store.Room().FindByID(ctx, updReq.RoomID)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)

	updReq.update(roomToUpdate)

	ownerPortalCode, assigneePortalCode := uint64(req.GetPortalCodeOwner()), uint64(req.GetPortalCodeAssignee())
	if _, err = s.roomService.Update(
		ctx,
		*user,
		*roomToUpdate,
		ownerPortalCode,
		assigneePortalCode,
		equipmentIDs,
		slotIDs,
		bookingIDs,
	); err != nil {
		return nil, err
	}

	pbRoom, err := s.roomFromDB(ctx, roomToUpdate.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateRoomResponse{
		Room: pbRoom,
	}, nil
}

func validateUpdateRequest(req *bookingpb.UpdateRoomRequest) (*updateRequest, []uuid.UUID, []uuid.UUID, []uuid.UUID, error) {
	var bookingIDs []uuid.UUID

	roomID, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, nil, nil, nil, status.New(codes.InvalidArgument, "invalid room_id").Err()
	}

	room, equipmentIDs, slotIDs, err := validateCreateDraftRequest(req)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if req.GetBookingIDs() != nil && len(req.GetBookingIDs()) > 0 {
		bookingIDs, err = getIDs(req.GetBookingIDs())
		if err != nil {
			return nil, nil, nil, nil, status.New(codes.InvalidArgument, "invalid booking id").Err()
		}
	}

	return &updateRequest{
		RoomID:      roomID,
		TargetAudience: room.TargetAudience,
		Title:          room.Title,
		Description:    room.Description,
		CreationDate:   room.CreationDate,
	}, equipmentIDs, slotIDs, bookingIDs, nil
}

type updateRequest struct {
	RoomID      uuid.UUID
	Owner          *model.User
	Employee       *model.User
	Title          string
	Description    string
	TargetAudience string
	CreationDate   time.Time
}

func (r updateRequest) update(room *model.Room) {
	room.Owner = r.Owner
	room.Employee = r.Employee
	room.Title = r.Title
	room.Description = r.Description
	room.TargetAudience = r.TargetAudience
	room.CreationDate = r.CreationDate
}
