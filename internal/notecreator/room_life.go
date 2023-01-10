package notecreator

import (
	"context"

	"be/internal/datastore/address"
	"be/internal/model"
	"be/internal/notes"
	"be/internal/notes/lifenote"
)

//nolint:gocognit
func (n NoteCreator) createRoomLifeNote(ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	room *model.Room,
) {
	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	link, err := n.getLink(n.cfg.BotID.String(), n.cfg.IP, n.cfg.BotToken, address.RoomAddressType, room.ID)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error get room link for life notice")
	}

	notice, err := lifenote.RoomNotice(notifyType, room, ctxUser, link)
	if err != nil {
		n.logger.Error().Err(err).Msg("error get room notice")
		return
	}

	email := false

	whomArr := notes.GetWhomRoomRecipientWithEmail(notifyType, room, admin, ctxUser, email)

	for _, whom := range whomArr {
		settings, err := n.store.User().GetNoteSettingByUser(ctx, &whom.UUID)
		if err != nil {
			n.logger.Error().Err(err).Msg("error get user settings for life note")
			continue
		}

		if settings.LifeOn {
			channelID, err := n.store.Address().GetChannelFromLifeNote(ctx, &whom.UUID)
			if err != nil {
				n.logger.Error().Err(err).Msg("error get channel for life note")
				continue
			}

			note := model.LifeNote{
				Event:       notifyType,
				Status:      model.NotSend,
				ActorID:     ctxUser.ID,
				Object:      model.AddressTypeRoom,
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
