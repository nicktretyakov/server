package notecreator

import (
	"context"

	"github.com/pkg/errors"

	"be/internal/datastore/base"
	"be/internal/datastore/address"
	"be/internal/model"
	"be/internal/notes"
	"be/internal/notes/lifenote"
)

//nolint:exhaustive,funlen,gocognit,gocyclo,cyclop
func (n NoteCreator) createBookingLifeNote(
	ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	bookings []*model.Booking,
) {
	var (
		linksBooking map[*model.Booking]string
		notice       *lifenote.Notice
		err          error
	)

	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	linksBooking, err = n.getLinksBookings(bookings)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error get booking link for life notice")
	}

	switch notifyType {
	case model.NotSendReportNotify:
		notice, err = lifenote.BookingNotSendReportNotifyNotice(bookings[0], linksBooking[bookings[0]])
	case model.MissedReportNotify:
		notice, err = lifenote.BookingMissedReportNotifyNotice(linksBooking)
	default:
		notice, err = lifenote.BookingNotice(notifyType, bookings[0], ctxUser, linksBooking[bookings[0]])
	}

	if err != nil {
		n.logger.Error().Err(err).Msg("error get booking notice")
		return
	}

	email := false

	whomArr := notes.GetWhomBookingRecipients(notifyType, bookings[0], admin, ctxUser, email)

	for _, whom := range whomArr {
		settings, err := n.store.User().GetNoteSettingByUser(ctx, &whom.UUID)
		if err != nil {
			n.logger.Error().Err(err).Msg("error get user settings for life note")
			continue
		}

		if settings.LifeOn {
			channelID, err := n.store.Address().GetChannelFromLifeNote(ctx, &whom.UUID)
			if err != nil && !errors.Is(err, base.ErrNotFound) {
				n.logger.Error().Err(err).Msg("error get channel for life note")
				continue
			}

			note := model.LifeNote{
				Event:       notifyType,
				Status:      model.NotSend,
				ActorID:     ctxUser.ID,
				Object:      model.AddressTypeBooking,
				RecipientID: whom.UUID,
				Body:        notice.Body,
				ChannelID:   channelID,
				BotID:       n.cfg.BotID,
				ForEntities: notice.ForEntities,
			}

			if _, err = n.store.Address().CreateLifeNote(ctx, note); err != nil {
				n.logger.Error().Err(err).Msg("error create life note")
			}
		}
	}
}

func (n NoteCreator) getLinksBookings(bookings []*model.Booking) (map[*model.Booking]string, error) {
	linksBooking := make(map[*model.Booking]string)

	for _, booking := range bookings {
		link, err := n.getLink(n.cfg.BotID.String(), n.cfg.IP, n.cfg.BotToken, address.BookingAddressType, booking.ID)
		if err != nil {
			return nil, err
		}

		linksBooking[booking] = link
	}

	return linksBooking, nil
}
