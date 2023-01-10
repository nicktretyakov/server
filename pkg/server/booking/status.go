package booking

import (
	"context"
	"time"

	bookingpb "be/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) ChangeBookingStatus(ctx context.Context, req *bookingpb.ChangeBookingStatusRequest) (*bookingpb.GetBookingResponse, error) { //nolint:lll
	user := auth.FromContext(ctx)

	bookingID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid booking_id").Err()
	}

	book, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	newStatus := model.Status(req.Status)
	if !acl.CanChangeStatus(*user, newStatus, book) {
		return nil, status.New(codes.PermissionDenied, "cannot change status").Err()
	}

	book.Status = newStatus

	if err = s.store.Booking().UpdateStatus(ctx, bookingID, newStatus); err != nil {
		return nil, err
	}

	if s.notificator != nil {
		go s.notificator.CreateBookingNote(user, getNoteTypeByStatus(newStatus), []*model.Booking{book}, time.Time{})
	}

	return &bookingpb.GetBookingResponse{
		Booking: pbs.PbBooking(*book),
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
