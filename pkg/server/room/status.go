package room

import (
	"context"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) ChangeRoomStatus(ctx context.Context, req *bookingpb.ChangeRoomStatusRequest) (*bookingpb.GetRoomResponse, error) {
	user := auth.FromContext(ctx)

	roomID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid room_id").Err()
	}

	room, err := s.store.Room().FindByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	newStatus := model.Status(req.Status)
	if !acl.CanChangeStatus(*user, newStatus, room) {
		return nil, status.New(codes.PermissionDenied, "cannot change status").Err()
	}

	room.Status = newStatus

	if err = s.store.Room().UpdateStatus(ctx, roomID, newStatus); err != nil {
		return nil, err
	}

	if s.notificator != nil {
		go s.notificator.CreateRoomNote(user, getNoteTypeByStatus(newStatus), room)
	}

	return &bookingpb.GetRoomResponse{
		Room: pbs.PbRoom(*room),
	}, nil
}

//nolint:exhaustive
func getNoteTypeByStatus(toStatus model.Status) model.NoteEvent {
	switch toStatus {
	case model.DeclinedAddressStatus: // На доработку
		return model.DeclinedNotify
	case model.OnRegisterAddressStatus: // На регистрации
		return model.OnRegisterNotify
	case model.ConfirmedAddressStatus: // Согласовано
		return model.ConfirmedNotify
	case model.DoneAddressStatus: // Завершен
		return model.DoneNotify
	case model.OnAgreeAddressStatus: // На согласовании
		return model.OnAgreeNotify
	default:
		return model.UnknownNotify
	}
}
