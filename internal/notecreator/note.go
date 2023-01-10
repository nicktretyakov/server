package notecreator

import (
	"context"
	"time"

	"be/internal/model"
)

func (n NoteCreator) CreateBookingNote(ctxUser *model.User,
	notifyType model.NoteEvent,
	bookings []*model.Booking,
	reportPeriod time.Time,
) {
	if n.cfg.EnabledNotify {
		n.createBookingEmailNote(context.Background(), ctxUser, notifyType, bookings, reportPeriod)
		n.createBookingLifeNote(context.Background(), ctxUser, notifyType, bookings)
		n.createBookingSystemNote(context.Background(), ctxUser, notifyType, bookings)
	}
}

func (n NoteCreator) CreateRoomNote(ctxUser *model.User, notifyType model.NoteEvent, room *model.Room) {
	if n.cfg.EnabledNotify {
		n.createRoomEmailNote(context.Background(), ctxUser, notifyType, room)
		n.createRoomLifeNote(context.Background(), ctxUser, notifyType, room)
		n.createRoomSystemNote(context.Background(), ctxUser, notifyType, room)
	}
}
