package notecreator

import (
	"context"

	"be/internal/model"
	"be/internal/notes"
	emailnote2 "be/internal/notes/emailnote"
)

func (n NoteCreator) createRoomEmailNote(
	ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	room *model.Room,
) {
	var (
		notice *emailnote2.Notice
		err    error
	)

	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	notice, err = emailnote2.RoomNotice(notifyType, room, ctxUser,
		notes.Link{
			Base:             n.cfg.URLBase,
			TypeAddress: n.cfg.RoomViewFrontRoute,
		})
	if err != nil {
		n.logger.Error().Err(err).Msg("error get room notice")
		return
	}

	email := true

	whomArr := notes.GetWhomRoomRecipientWithEmail(notifyType, room, admin, ctxUser, email)

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
				Object:         model.AddressTypeRoom,
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
