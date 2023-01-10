package lifenote

import (
	"github.com/pkg/errors"

	"be/internal/model"
	"be/internal/notes"
)

type Notice struct {
	Body        string
	ForEntities []model.ForEntity
}

func BookingMissedReportNotifyNotice(bookings map[*model.Booking]string) (*Notice, error) {
	dataForBody := make([]*notes.MissedReportNotifyData, 0, len(bookings))

	for booking, shortLink := range bookings {
		periods := make([]string, 0, len(booking.Reports))

		for _, report := range booking.Reports {
			if report.NotRelevant() {
				periods = append(periods, notes.YearMonthString(report.Period.Time()))
			}
		}

		dataForBody = append(dataForBody, &notes.MissedReportNotifyData{
			BookingTitle: booking.Title,
			BookingLink:  shortLink,
			Periods:      periods,
			Rpo:          booking.Supervisor,
		})
	}

	return getMissedReportNotify(dataForBody), nil
}

func BookingNotSendReportNotifyNotice(booking *model.Booking, shortLink string) (*Notice, error) {
	periods := make([]string, 0, len(booking.Reports))

	for _, report := range booking.Reports {
		if report.NotRelevant() {
			periods = append(periods, notes.YearMonthString(report.Period.Time()))
		}
	}

	return getNotSendReportNotify(booking.Title, shortLink, periods), nil
}

//nolint:cyclop,exhaustive,ineffassign,wastedassign,staticcheck
func BookingNotice(notify model.NoteEvent,
	booking *model.Booking,
	actor *model.User,
	shortLink string,
) (*Notice, error) {
	data := &Notice{}

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(booking.Title, notes.AddressTypeBooking, shortLink)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(booking.Title, notes.AddressTypeBooking, shortLink, booking.Assignee)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(booking.Title, notes.AddressTypeBooking, shortLink, actor)
	case model.DeclinedNotify:
		data = getDeclinedNotify(booking.Title, notes.AddressTypeBooking, shortLink, actor)
	case model.DoneNotify:
		data = getDoneNotify(booking.Title, shortLink)
	case model.SentReportNotify:
		data = getSentReportNotify(booking.Title, shortLink, actor)
	case model.FinalReportOnRegisterNotify:
		data = getFinalReportOnRegisterNotify(booking.Title, shortLink)
	case model.FinalReportOnAgreeNotify:
		data = getFinalReportOnAgreeNotify(booking.Title, shortLink, booking.Assignee)
	case model.FinalReportDeclinedNotify:
		data = getFinalReportDeclinedNotify(booking.Title, shortLink, actor)
	default:
		return nil, errors.New("unexpected notify type for booking")
	}

	return data, nil
}

//nolint:exhaustive,ineffassign,wastedassign,staticcheck
func RoomNotice(notify model.NoteEvent, room *model.Room, actor *model.User, shortLink string) (*Notice, error) {
	data := &Notice{}

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(room.Title, notes.AddressTypeRoom, shortLink)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(room.Title, notes.AddressTypeRoom, shortLink, room.Employee)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(room.Title, notes.AddressTypeRoom, shortLink, actor)
	case model.DeclinedNotify:
		data = getDeclinedNotify(room.Title, notes.AddressTypeRoom, shortLink, actor)
	default:
		return nil, errors.New("unexpected notify type for room")
	}

	return data, nil
}
