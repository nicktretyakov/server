package outmember

import (
	"context"
	"time"
	"unicode/utf8"

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

const minCommentLen = 3

func (s Service) AcceptBookingAsRegister(ctx context.Context, req *bookingpb.BookingRequest) (*bookingpb.BookingResponse, error) {
	return s.newInitialReport(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineBookingAsRegister(ctx context.Context, req *bookingpb.BookingRequest) (*bookingpb.BookingResponse, error) {
	return s.newInitialReport(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) AcceptBookingAsAssignee(ctx context.Context, req *bookingpb.BookingRequest) (*bookingpb.BookingResponse, error) {
	return s.newApprovalReport(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineBookingAsAssignee(ctx context.Context, req *bookingpb.BookingRequest) (*bookingpb.BookingResponse, error) {
	return s.newApprovalReport(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) newInitialReport(ctx context.Context, req *bookingpb.BookingRequest,
	result model.OutmemberResult,
) (*bookingpb.BookingResponse, error) {
	bookingID, err := uuid.Parse(req.GetBookingID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid booking_id").Err()
	}

	user := auth.FromContext(ctx)

	if !acl.IsHeadOfBooking(*user) {
		return nil, status.New(codes.PermissionDenied, "superuser only").Err()
	}

	if result == model.DeclineOutmemberResult && utf8.RuneCountInString(req.GetComment()) < minCommentLen {
		return nil, status.New(codes.InvalidArgument, "invalid comment").Err()
	}

	if err = s.outmemberService.NewInitialOutmember(ctx, *user, addressStore.BookingAddressType, model.Outmember{
		Type:           model.InitialOutmemberType,
		UserID:         user.ID,
		AddressID: bookingID,
		Result:         result,
		Extra:          map[string]interface{}{"comment": req.GetComment()},
		Role:           model.BookingManagerOutmemberRole,
	}); err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateBookingNote(user,
		getNoteTypeByStatus(booking.Status),
		[]*model.Booking{booking},
		time.Time{})

	return &bookingpb.BookingResponse{Booking: pbs.PbBooking(*booking)}, nil
}

func (s Service) newApprovalReport(ctx context.Context, req *bookingpb.BookingRequest,
	result model.OutmemberResult,
) (*bookingpb.BookingResponse, error) {
	bookingID, err := uuid.Parse(req.GetBookingID())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid booking_id").Err()
	}

	user := auth.FromContext(ctx)

	if result == model.DeclineOutmemberResult && utf8.RuneCountInString(req.GetComment()) < minCommentLen {
		return nil, status.New(codes.InvalidArgument, "invalid comment").Err()
	}

	if err = s.outmemberService.NewAcceptanceOutmember(ctx, *user, addressStore.BookingAddressType, model.Outmember{
		Type:           model.ApprovalOutmemberType,
		UserID:         user.ID,
		AddressID: bookingID,
		Result:         result,
		Extra:          map[string]interface{}{"comment": req.GetComment()},
		Role:           model.AssigneeOutmemberRole,
	}); err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateBookingNote(user,
		getNoteTypeByStatus(booking.Status),
		[]*model.Booking{booking},
		time.Time{})

	return &bookingpb.BookingResponse{Booking: pbs.PbBooking(*booking)}, nil
}
