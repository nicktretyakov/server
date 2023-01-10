package outmember

import (
	"context"
	"time"

	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/acl"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/pbs"
)

func (s Service) AcceptFinalReportAsAssignee(ctx context.Context,
	req *bookingpb.FinalReportRequest,
) (*bookingpb.FinalReportResponse, error) {
	return s.acceptFinalReport(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineFinalReportAsAssignee(ctx context.Context,
	req *bookingpb.FinalReportRequest,
) (*bookingpb.FinalReportResponse, error) {
	return s.acceptFinalReport(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) AcceptFinalReportAsRegister(ctx context.Context,
	req *bookingpb.FinalReportRequest,
) (*bookingpb.FinalReportResponse, error) {
	return s.newFinalReport(ctx, req, model.AcceptOutmemberResult)
}

func (s Service) DeclineFinalReportAsRegister(ctx context.Context,
	req *bookingpb.FinalReportRequest,
) (*bookingpb.FinalReportResponse, error) {
	return s.newFinalReport(ctx, req, model.DeclineOutmemberResult)
}

func (s Service) newFinalReport(ctx context.Context, req *bookingpb.FinalReportRequest,
	result model.OutmemberResult,
) (*bookingpb.FinalReportResponse, error) {
	finalReportID, err := uuid.Parse(req.GetFinalReportId())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid final_report_id").Err()
	}

	user := auth.FromContext(ctx)

	if !acl.IsHeadOfBooking(*user) {
		return nil, status.New(codes.PermissionDenied, "superuser only").Err()
	}

	booking, err := s.outmemberService.RegisterFinalOutmember(ctx, *user, finalReportID, model.Outmember{
		Type:   model.FinalRegisterOutmemberType,
		Result: result,
		Extra:  map[string]interface{}{"comment": req.GetComment()},
		Role:   model.BookingManagerOutmemberRole,
	})
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateBookingNote(user,
		getNotificationTypeByFinalStatus(booking.FinalReport.Status),
		[]*model.Booking{booking},
		time.Time{})

	return &bookingpb.FinalReportResponse{
		Role:    bookingpb.Role_ROLE_BOOKING_MANAGER,
		Booking: pbs.PbBooking(*booking),
		User:    pbs.PbUser(user),
	}, nil
}

func (s Service) acceptFinalReport(ctx context.Context, req *bookingpb.FinalReportRequest,
	result model.OutmemberResult,
) (*bookingpb.FinalReportResponse, error) {
	finalReportID, err := uuid.Parse(req.GetFinalReportId())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid final_report_id").Err()
	}

	user := auth.FromContext(ctx)

	booking, err := s.outmemberService.AcceptFinalOutmember(ctx, *user, finalReportID, model.Outmember{
		Type:   model.FinalApprovalOutmemberType,
		Result: result,
		Extra:  map[string]interface{}{"comment": req.GetComment()},
		Role:   model.AssigneeOutmemberRole,
	})
	if err != nil {
		return nil, err
	}

	go s.Notificator.CreateBookingNote(user,
		getNotificationTypeByFinalStatus(booking.FinalReport.Status),
		[]*model.Booking{booking},
		time.Time{})

	if booking.FinalReport.Status == model.ConfirmedFinalReportStatus {
		go s.Notificator.CreateBookingNote(user,
			getNoteTypeByStatus(booking.Status),
			[]*model.Booking{booking},
			time.Time{})
	}

	return &bookingpb.FinalReportResponse{
		Role:    bookingpb.Role_ROLE_ASSIGNEE,
		Booking: pbs.PbBooking(*booking),
		User:    pbs.PbUser(user),
	}, nil
}

//nolint:exhaustive
func getNotificationTypeByFinalStatus(toStatus model.FinalReportStatus) model.NoteEvent {
	switch toStatus {
	case model.OnRegisterFinalReportStatus:
		return model.FinalReportOnRegisterNotify
	case model.OnAgreeFinalReportStatus:
		return model.FinalReportOnAgreeNotify
	case model.DeclinedFinalReportStatus:
		return model.FinalReportDeclinedNotify
	default:
		return model.UnknownNotify
	}
}
