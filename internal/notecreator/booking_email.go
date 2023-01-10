package notecreator

import (
	"context"
	"time"

	"be/internal/model"
	"be/internal/notes"
	emailnote2 "be/internal/notes/emailnote"
)

//nolint:exhaustive, funlen, gocognit
func (n NoteCreator) createBookingEmailNote(
	ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	bookings []*model.Booking,
	reportPeriod time.Time,
) {
	var (
		notice *emailnote2.Notice
		err    error
	)

	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	switch notifyType {
	case model.NotSendReportNotify:
		notice, err = emailnote2.BookingNotSendReportNotifyNotice(bookings[0],
			notes.Link{
				Base:             n.cfg.URLBase,
				TypeAddress: n.cfg.BookingViewFrontRoute,
			})
	case model.MissedReportNotify:
		notice, err = emailnote2.BookingMissedReportNotifyNotice(bookings, notes.Link{
			Base:             n.cfg.URLBase,
			TypeAddress: n.cfg.BookingViewFrontRoute,
		})
	default:
		notice, err = emailnote2.BookingNotice(notifyType, bookings[0], ctxUser, reportPeriod, notes.Link{
			Base:             n.cfg.URLBase,
			TypeAddress: n.cfg.BookingViewFrontRoute,
		})
	}

	if err != nil {
		n.logger.Error().Err(err).Msg("error get booking notice")
		return
	}

	email := true

	whomArr := notes.GetWhomBookingRecipients(notifyType, bookings[0], admin, ctxUser, email)

	for _, whom := range whomArr {
		settings, err := n.store.User().GetNoteSettingByUser(ctx, &whom.UUID)
		if err != nil {
			n.logger.Error().Err(err).Msg("error get user settings for email note")
			continue
		}

		if settings.EmailOn {
			note := model.EmailNote{
				Event:          notifyType,
				Status:         model.NotSend,
				ActorID:        ctxUser.ID,
				Object:         model.AddressTypeBooking,
				RecipientID:    whom.UUID,
				RecipientEmail: whom.Email,
				SenderEmail:    n.cfg.Sender,
				Subject:        notice.Subject,
				Body:           notice.Body,
			}

			if _, err := n.store.Address().CreateEmailNote(ctx, note); err != nil {
				n.logger.Error().Err(err).Msg("error create email note")
			}
		}
	}
}
