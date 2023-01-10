package finalreport

import (
	"context"
	"time"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/lib"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
	"be/pkg/server/validators"
)

func (s Service) SendReport(ctx context.Context, req *bookingpb.SendFinalReportRequest) (*bookingpb.SendFinalReportResponse, error) {
	bookingID, err := uuid.Parse(req.GetBookingId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid booking_id")
	}

	var user *model.User

	if user = auth.FromContext(ctx); user == nil {
		return nil, status.Error(codes.Unauthenticated, "auth required")
	}

	slot, err := validators.Notification(req.GetSlot())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "slot invalid: %s", err.Error())
	}

	endAt, err := validators.Time(req.GetEndDate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "time invalid: %s", err.Error())
	}

	attachmentsUUID, err := lib.ParseUUIDStrings(req.GetAttachmentId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "attachment uuid invalid: %s", err.Error())
	}

	rep, err := s.bookingService.SendFinalReport(ctx, *user, model.FinalReport{
		BookingID:       bookingID,
		Slot:          slot,
		EndAt:           endAt,
		Comment:         req.GetComment(),
		AttachmentsUUID: attachmentsUUID,
	})
	if err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	go s.notificator.CreateBookingNote(user,
		model.FinalReportOnRegisterNotify,
		[]*model.Booking{booking},
		time.Time{})

	return &bookingpb.SendFinalReportResponse{
		Report:        pbs.FinalReport(*rep, booking.GetAttachmentsFinalReport()),
		BookingStatus: bookingpb.AddressStatus(booking.Status),
	}, nil
}
