package systemnote

import (
	"errors"
	"fmt"

	"be/internal/model"
	"be/internal/notes"
)

type Notice struct {
	Header string
	Body   string
}

func BookingMissedReportNotifyNotice(bookings []*model.Booking, link notes.Link) (*Notice, error) {
	dataForBody := make([]*notes.MissedReportNotifyData, 0, len(bookings))

	for _, booking := range bookings {
		periods := make([]string, 0, len(booking.Reports))

		for _, report := range booking.Reports {
			if report.NotRelevant() {
				periods = append(periods, notes.YearMonthString(report.Period.Time()))
			}
		}

		dataForBody = append(dataForBody, &notes.MissedReportNotifyData{
			BookingTitle: booking.Title,
			BookingLink: fmt.Sprintf("<a href=\"%s\"> %s </a>",
				notes.CreateAddressViewLink(booking.ID, link),
				booking.Title),
			Periods: periods,
			Rpo:     booking.Supervisor,
		})
	}

	return getMissedReportNotify(dataForBody), nil
}

func BookingNotSendReportNotifyNotice(booking *model.Booking) (*Notice, error) {
	periods := make([]string, 0, len(booking.Reports))

	for _, report := range booking.Reports {
		if report.NotRelevant() {
			periods = append(periods, notes.YearMonthString(report.Period.Time()))
		}
	}

	return getNotSendReportNotify(booking.Title, periods), nil
}

//nolint:cyclop,exhaustive,ineffassign,wastedassign,staticcheck
func BookingNotice(notify model.NoteEvent,
	booking *model.Booking,
	actor *model.User,
) (*Notice, error) {
	data := &Notice{}

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(booking.Title, notes.AddressTypeBooking)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(booking.Title, notes.AddressTypeBooking, booking.Assignee)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(booking.Title, notes.AddressTypeBooking, actor)
	case model.DeclinedNotify:
		data = getDeclinedNotify(booking.Title, notes.AddressTypeBooking, actor)
	case model.DoneNotify:
		data = getDoneNotify(booking.Title)
	case model.SentReportNotify:
		data = getSentReportNotify(booking.Title, actor)
	case model.FinalReportOnRegisterNotify:
		data = getFinalReportOnRegisterNotify(booking.Title)
	case model.FinalReportOnAgreeNotify:
		data = getFinalReportOnAgreeNotify(booking.Title, booking.Assignee)
	case model.FinalReportDeclinedNotify:
		data = getFinalReportDeclinedNotify(booking.Title, actor)
	default:
		return nil, errors.New("unexpected notify type for booking")
	}

	return data, nil
}

//nolint:exhaustive,ineffassign,wastedassign,staticcheck
func RoomNotice(notify model.NoteEvent, room *model.Room, actor *model.User) (*Notice, error) {
	data := &Notice{}

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(room.Title, notes.AddressTypeRoom)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(room.Title, notes.AddressTypeRoom, room.Employee)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(room.Title, notes.AddressTypeRoom, actor)
	case model.DeclinedNotify:
		data = getDeclinedNotify(room.Title, notes.AddressTypeRoom, actor)
	default:
		return nil, errors.New("unexpected notify type for room")
	}

	return data, nil
}
