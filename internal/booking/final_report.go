package booking

import (
	"context"

	"github.com/pkg/errors"

	"be/internal/acl"
	"be/internal/datastore/base"
	"be/internal/model"
)

func (s Service) SendFinalReport(ctx context.Context, user model.User, finalReport model.FinalReport) (*model.FinalReport, error) {
	finalReport.Status = model.InitialFinalReportStatus

	existedReport, err := s.store.FinalReport().FindByBookingID(ctx, finalReport.BookingID)
	if err != nil && !errors.Is(err, base.ErrNotFound) {
		return nil, err
	}

	if existedReport != nil {
		finalReport.ID = existedReport.ID
		finalReport.CreatedAt = existedReport.CreatedAt
		finalReport.UpdatedAt = existedReport.UpdatedAt
		finalReport.Status = existedReport.Status
	}

	book, err := s.store.Booking().FindByID(ctx, finalReport.BookingID)
	if err != nil {
		return nil, err
	}

	if !acl.CanSendFinalReport(user, *book) {
		return nil, ErrPermissionDenied
	}

	if err = s.store.Booking().UpdateStatus(ctx, book.ID, model.FinalizeOnRegisterStatus); err != nil {
		return nil, err
	}

	finalReport.Status = model.OnRegisterFinalReportStatus

	if existedReport != nil {
		return s.store.FinalReport().Update(ctx, finalReport)
	}

	return s.store.FinalReport().Create(ctx, finalReport)
}
