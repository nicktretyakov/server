package notecreator

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
	"be/internal/notes"
	"be/internal/notes/systemnote"
)

// nolint: exhaustive
func (n NoteCreator) createBookingSystemNote(
	ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	bookings []*model.Booking,
) {
	var (
		notice    *systemnote.Notice
		bookingID uuid.UUID
	)

	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	switch notifyType {
	case model.NotSendReportNotify:
		notice, err = systemnote.BookingNotSendReportNotifyNotice(bookings[0])

		bookingID = bookings[0].ID
	case model.MissedReportNotify:
		notice, err = systemnote.BookingMissedReportNotifyNotice(bookings,
			notes.Link{
				Base:             n.cfg.URLBase,
				TypeAddress: n.cfg.BookingViewFrontRoute,
			})

		bookingID = uuid.Nil
	default:
		notice, err = systemnote.BookingNotice(notifyType, bookings[0], ctxUser)

		bookingID = bookings[0].ID
	}

	if err != nil {
		n.logger.Error().Err(err).Msg("error get booking notice")
		return
	}

	email := false

	whomArr := notes.GetWhomBookingRecipients(notifyType, bookings[0], admin, ctxUser, email)

	for _, whom := range whomArr {
		note := model.SystemNote{
			Event:          notifyType,
			Status:         model.NotRead,
			ActorID:        ctxUser.ID,
			AddressID: &bookingID,
			Object:         model.AddressTypeBooking,
			RecipientID:    whom.UUID,
			Header:         notice.Header,
			Body:           notice.Body,
		}

		if _, err = n.store.Address().CreateSystemNote(ctx, note); err != nil {
			n.logger.Error().Err(err).Msg("error create life note")
		}
	}
}
