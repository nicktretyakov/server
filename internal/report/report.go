package report

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
)

func (s service) CheckReports(ctx context.Context, bookingID *uuid.UUID, periods []model.Period) error {
	if bookingID == nil {
		return nil
	}

	reports, err := s.store.ListByBookingID(ctx, *bookingID)
	if err != nil {
		return err
	}

	if len(reports) == 0 {
		if err = s.createNewReports(ctx, *bookingID, periods); err != nil {
			return err
		}

		return nil
	}

	if err = s.deleteReportsNotSend(ctx, reports); err != nil {
		return err
	}

	if err = s.createReportsByNewPeriod(ctx, *bookingID, periods, reports); err != nil {
		return err
	}

	return nil
}

func (s service) deleteReportsNotSend(ctx context.Context, reports []model.ReportBooking) error {
	deleteReport := make([]model.ReportBooking, 0, len(reports))

	for _, report := range reports {
		if report.Status == model.NotSendReportStatus {
			deleteReport = append(deleteReport, report)
		}
	}

	if len(deleteReport) == 0 {
		return nil
	}

	if err := s.store.BulkDelete(ctx, deleteReport); err != nil {
		return err
	}

	return nil
}

func (s service) createReportsByNewPeriod(
	ctx context.Context,
	bookingID uuid.UUID,
	periods []model.Period,
	reports []model.ReportBooking,
) error {
	createdReport := make([]model.ReportBooking, 0, len(reports))

	for _, period := range periods {
		report := findReportByPeriod(period, reports)

		if report != nil {
			if report.Status == model.SendReportStatus {
				continue
			}

			createdReport = append(createdReport, *createReportModel(bookingID, report.Period))
		} else {
			createdReport = append(createdReport, *createReportModel(bookingID, period))
		}
	}

	if len(createdReport) == 0 {
		return nil
	}

	if _, err := s.store.BulkCreate(ctx, createdReport); err != nil {
		return err
	}

	return nil
}

func findReportByPeriod(period model.Period, reports []model.ReportBooking) *model.ReportBooking {
	for _, report := range reports {
		if period.Month == report.Period.Month && period.Year == report.Period.Year {
			return &report
		}
	}

	return nil
}

func (s service) createNewReports(ctx context.Context, bookingID uuid.UUID, periods []model.Period) error {
	if len(periods) == 0 {
		return nil
	}

	reports := make([]model.ReportBooking, 0, len(periods))

	for _, period := range periods {
		reports = append(reports, *createReportModel(bookingID, period))
	}

	if _, err := s.store.BulkCreate(ctx, reports); err != nil {
		return err
	}

	return nil
}

func createReportModel(bookingID uuid.UUID, period model.Period) *model.ReportBooking {
	return &model.ReportBooking{
		BookingID: bookingID,
		Period:    period,
		Status:    model.NotSendReportStatus,
	}
}
