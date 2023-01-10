package notecreator

import (
	"context"

	"be/internal/model"
	"be/internal/notes"
	"be/internal/notes/systemnote"
)

func (n NoteCreator) createRoomSystemNote(
	ctx context.Context,
	ctxUser *model.User,
	notifyType model.NoteEvent,
	room *model.Room,
) {
	admin, err := n.store.User().FindAdmins(ctx)
	if err != nil {
		n.logger.Warn().Err(err).Msg("error find admin")
	}

	notice, err := systemnote.RoomNotice(notifyType, room, ctxUser)
	if err != nil {
		n.logger.Error().Err(err).Msg("error get room notice")
		return
	}

	email := false

	whomArr := notes.GetWhomRoomRecipientWithEmail(notifyType, room, admin, ctxUser, email)

	for _, whom := range whomArr {
		note := model.SystemNote{
			Event:          notifyType,
			Status:         model.NotRead,
			ActorID:        ctxUser.ID,
			AddressID: &room.ID,
			Object:         model.AddressTypeRoom,
			RecipientID:    whom.UUID,
			Header:         notice.Header,
			Body:           notice.Body,
		}

		if _, err = n.store.Address().CreateSystemNote(ctx, note); err != nil {
			n.logger.Error().Err(err).Msg("error create life note")
		}
	}
}
