package emailnote

import (
	"time"

	"github.com/pkg/errors"

	"be/internal/model"
	"be/internal/notes"
)

type Notice struct {
	Subject string
	Body    string
}

func BookingMissedReportNotifyNotice(bookings []*model.Booking, link notes.Link) (*Notice, error) {
	dataForBody := make([]*MissedReportNotifyData, 0, len(bookings))

	for _, booking := range bookings {
		periods := make([]string, 0, len(booking.Reports))

		for _, report := range booking.Reports {
			if report.NotRelevant() {
				periods = append(periods, notes.YearMonthString(report.Period.Time()))
			}
		}

		dataForBody = append(dataForBody, &MissedReportNotifyData{
			BookingTitle: booking.Title,
			BookingLink:  notes.CreateAddressViewLink(booking.ID, link),
			Periods:      periods,
			Rpo:          booking.Supervisor,
		})
	}

	subject := getSubject(model.MissedReportNotify, "", notes.AddressTypeBooking, "")
	data := getMissedReportNotify(dataForBody)

	body, err := executeBody(subject, data)
	if err != nil {
		return nil, err
	}

	return &Notice{
		Subject: subject,
		Body:    body,
	}, nil
}

func BookingNotSendReportNotifyNotice(booking *model.Booking, link notes.Link) (*Notice, error) {
	periods := make([]string, 0, len(booking.Reports))

	for _, report := range booking.Reports {
		if report.NotRelevant() {
			periods = append(periods, notes.YearMonthString(report.Period.Time()))
		}
	}

	subject := getSubject(model.NotSendReportNotify, booking.Title, notes.AddressTypeBooking, "")
	data := getNotSendReportNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link), periods)

	body, err := executeBody(subject, data)
	if err != nil {
		return nil, err
	}

	return &Notice{
		Subject: subject,
		Body:    body,
	}, nil
}

//nolint:gocyclo,cyclop,exhaustive,ineffassign,wastedassign,funlen
func BookingNotice(notify model.NoteEvent,
	booking *model.Booking,
	actor *model.User,
	reportPeriod time.Time,
	link notes.Link,
) (*Notice, error) {
	var (
		subject = getSubject(notify, booking.Title, notes.AddressTypeBooking, notes.YearMonthString(reportPeriod))
		data    = &Data{}
	)

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(
			booking.Title,
			notes.AddressTypeBooking,
			notes.CreateAddressViewLink(booking.ID, link),
		)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(
			booking.Title,
			notes.AddressTypeBooking,
			notes.CreateAddressViewLink(booking.ID, link),
			booking.Assignee,
		)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(
			booking.Title,
			notes.AddressTypeBooking,
			notes.CreateAddressViewLink(booking.ID, link),
			actor,
		)
	case model.DeclinedNotify:
		data = getDeclinedNotify(
			booking.Title,
			notes.AddressTypeBooking,
			notes.CreateAddressViewLink(booking.ID, link),
			actor,
		)
	case model.DoneNotify:
		data = getDoneNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link), actor)
	case model.SentReportNotify:
		data = getSentReportNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link), actor)
	case model.FinalReportOnRegisterNotify:
		data = getFinalReportOnRegisterNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link))
	case model.FinalReportOnAgreeNotify:
		data = getFinalReportOnAgreeNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link), booking.Assignee)
	case model.FinalReportDeclinedNotify:
		data = getFinalReportDeclinedNotify(booking.Title, notes.CreateAddressViewLink(booking.ID, link), actor)
	default:
		return nil, errors.New("unexpected notify type for booking")
	}

	body, err := executeBody(subject, data)
	if err != nil {
		return nil, err
	}

	return &Notice{
		Subject: subject,
		Body:    body,
	}, nil
}

//nolint:exhaustive,ineffassign,wastedassign
func RoomNotice(notify model.NoteEvent, room *model.Room, actor *model.User, link notes.Link) (*Notice, error) {
	var (
		subject = getSubject(notify, room.Title, notes.AddressTypeRoom, "")
		data    = &Data{}
	)

	switch notify {
	case model.OnRegisterNotify:
		data = getOnRegisterNotify(
			room.Title,
			notes.AddressTypeRoom,
			notes.CreateAddressViewLink(room.ID, link),
		)
	case model.OnAgreeNotify:
		data = getOnAgreeNotify(
			room.Title,
			notes.AddressTypeRoom,
			notes.CreateAddressViewLink(room.ID, link),
			room.Employee,
		)
	case model.ConfirmedNotify:
		data = getConfirmedNotify(
			room.Title,
			notes.AddressTypeRoom,
			notes.CreateAddressViewLink(room.ID, link),
			actor,
		)
	case model.DeclinedNotify:
		data = getDeclinedNotify(
			room.Title,
			notes.AddressTypeRoom,
			notes.CreateAddressViewLink(room.ID, link),
			actor,
		)
	default:
		return nil, errors.New("unexpected notify type for room")
	}

	body, err := executeBody(subject, data)
	if err != nil {
		return nil, err
	}

	return &Notice{
		Subject: subject,
		Body:    body,
	}, nil
}
