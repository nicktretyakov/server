package booking

import (
	"context"

	"be/internal/acl"
	"be/internal/model"
)

func (s Service) SendReport(ctx context.Context, user model.User, updReport model.ReportBooking) (*model.ReportBooking, error) {
	rep, err := s.store.Report().FindBookingReportByID(ctx, updReport.ID)
	if err != nil {
		return nil, err
	}

	book, err := s.store.Booking().FindByID(ctx, rep.BookingID)
	if err != nil {
		return nil, err
	}

	if !acl.CanSendReport(user, *book, *rep) {
		return nil, ErrPermissionDenied
	}

	rep.Slot = updReport.Slot
	rep.EndAt = updReport.EndAt
	rep.Events = updReport.Events
	rep.Reasons = updReport.Reasons
	rep.Comment = updReport.Comment
	rep.Status = model.SendReportStatus

	return s.store.Report().Update(ctx, *rep)
}
