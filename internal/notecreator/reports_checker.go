package notecreator

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"

	"be/internal/model"
)

func (n *NoteCreator) reportsChecker(ctx context.Context) {
	moscowTime, err := time.LoadLocation("Local")
	if err != nil {
		n.logger.Error().Err(err).Msg("Load Location failed")
		return
	}

	n.reportsCheckerScheduler = cron.New(cron.WithLocation(moscowTime))

	if _, err := n.reportsCheckerScheduler.AddFunc(n.cfg.ReportsChecker.getSchedulerString(), func() {
		n.reportsCheckerEmailsCreate(ctx)
	}); err != nil {
		n.logger.Error().Err(err).Msg("Run emails create by reports checker scheduler failed")
		return
	}

	go n.reportsCheckerScheduler.Start()
}

func (n *NoteCreator) reportsCheckerEmailsCreate(ctx context.Context) {
	bookings, err := n.store.Booking().ListBookingsWithNotSendReports(ctx)
	if err != nil {
		n.logger.Error().Err(err).Msg("ListBookingWithNotSendReports return error")
		return
	}

	for i := range bookings {
		go func(booking *model.Booking) {
			for _, report := range booking.Reports {
				if report.Status == model.NotSendReportStatus {
					n.CreateBookingNote(&model.User{}, model.NotSendReportNotify, []*model.Booking{booking}, report.Period.Time())

					break
				}
			}
		}(&bookings[i])
	}
}
