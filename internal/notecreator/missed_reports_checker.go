package notecreator

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"

	"be/internal/model"
)

func (n *NoteCreator) missedReportsChecker(ctx context.Context) {
	moscowTime, err := time.LoadLocation("Local")
	if err != nil {
		n.logger.Error().Err(err).Msg("Load Location failed")
		return
	}

	n.missedReportsCheckerScheduler = cron.New(cron.WithLocation(moscowTime))

	if _, err := n.reportsCheckerScheduler.AddFunc(n.cfg.MissedReportsChecker.getSchedulerString(), func() {
		n.missedReportsCheckerEmailsCreate(ctx)
	}); err != nil {
		n.logger.Error().Err(err).Msg("Run emails create by reports checker scheduler failed")
		return
	}

	go n.reportsCheckerScheduler.Start()
}

func (n *NoteCreator) missedReportsCheckerEmailsCreate(ctx context.Context) {
	bookings, err := n.store.Booking().ListBookingsWithNotSendReports(ctx)
	if err != nil {
		n.logger.Error().Err(err).Msg("ListBookingWithNotSendReports return error")
		return
	}

	bookingsMap := make(map[uuid.UUID][]*model.Booking)
	bookingsNew := make([]*model.Booking, 0, len(bookings))

	for i, booking := range bookings {
		if booking.Supervisor == nil {
			continue
		}

		if _, ok := bookingsMap[booking.Supervisor.ID]; !ok {
			bookingsMap[booking.Supervisor.ID] = make([]*model.Booking, 0, 10)
		}

		bookingsMap[booking.Supervisor.ID] = append(bookingsMap[booking.Supervisor.ID], &bookings[i])
	}

	for _, v := range bookingsMap {
		bookingsNew = append(bookingsNew, v...)
	}

	n.CreateBookingNote(&model.User{}, model.MissedReportNotify, bookingsNew, time.Time{})
}
