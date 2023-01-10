package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func PbSystemNote(note *model.SystemNote) *bookingpb.SystemNote {
	status := bookingpb.STATUS_NOTE_READ

	if note.Status == model.NotRead {
		status = bookingpb.STATUS_NOTE_NOT_READ
	}

	object := bookingpb.AddressType_BOOKING

	if note.Object == model.AddressTypeRoom {
		object = bookingpb.AddressType_ROOM
	}

	return &bookingpb.SystemNote{
		Id:               note.ID.String(),
		Status:           status,
		AddressId:   note.AddressID.String(),
		AddressType: object,
		Header:           note.Header,
		Body:             note.Body,
		CreatedAt:        ToUTCString(note.CreatedAt),
		ReadAt:           ToUTCString(note.ReadAt),
	}
}

func PbSystemsNotes(notes []*model.SystemNote) []*bookingpb.SystemNote {
	pbSystemNotes := make([]*bookingpb.SystemNote, 0, len(notes))

	for _, note := range notes {
		pbSystemNotes = append(pbSystemNotes, PbSystemNote(note))
	}

	return pbSystemNotes
}
