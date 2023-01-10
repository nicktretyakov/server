package pbs

import (
	"github.com/google/uuid"
	bookingpb "be/proto"

	"be/internal/model"
)

func PbBooking(booking model.Booking) *bookingpb.Booking {
	return &bookingpb.Booking{
		Uuid:  booking.ID.String(),
		Type:  bookingpb.BookingType(booking.Type),
		Title: booking.Title,
		City:  booking.City,
		Timeline: &bookingpb.Timeline{
			Start: ToUTCString(booking.Timeline.StartAt),
			End:   ToUTCString(booking.Timeline.EndAt),
		},
		Description:  booking.Description,
		Department:   ToPbDepartments(booking.Departments),
		Slot:       toNotification(booking.Slot),
		Supervisor:   PbUser(booking.Supervisor),
		Assignee:     PbUser(booking.Assignee),
		Author:       PbUser(booking.Author),
		Goal:         booking.Goal,
		Status:       bookingpb.AddressStatus(booking.Status),
		CreatedAt:    ToUTCString(booking.CreatedAt),
		UpdatedAt:    ToUTCString(booking.UpdatedAt),
		Number:       booking.Number,
		Reports:      BookingReportsList(booking.Reports),
		FinalReport:  FinalReport(booking.FinalReport, booking.GetAttachmentsFinalReport()),
		Outmembers:   OutmembersList(booking.Outmembers),
		Attachments:  AttachmentList(booking.Attachments),
		Links:        PbLinkList(booking.Links),
		Participants: PbParticipantsList(booking.Participants),
		State:        bookingpb.AddressState(booking.State),
	}
}

func PbUser(user *model.User) *bookingpb.User {
	if user == nil {
		return nil
	}

	return &bookingpb.User{
		Uuid:      user.ID.String(),
		ProfileId: user.ProfileID,
		Email:     user.Email,
		Phone:     user.Phone,
		Employee:  PbEmployee(user.Employee),
	}
}

func PbUserList(list []model.User) []*bookingpb.User {
	res := make([]*bookingpb.User, 0, len(list))

	for i := range list {
		res = append(res, PbUser(&list[i]))
	}

	return res
}

func PbBookings(bookings []model.Booking) []*bookingpb.Booking {
	pbBookings := make([]*bookingpb.Booking, 0, len(bookings))
	for _, booking := range bookings {
		pbBookings = append(pbBookings, PbBooking(booking))
	}

	return pbBookings
}

func PbLinkList(links []model.Link) []*bookingpb.Link {
	pbLinks := make([]*bookingpb.Link, 0, len(links))
	for _, link := range links {
		pbLinks = append(pbLinks, &bookingpb.Link{
			Id:     link.Id.String(),
			Name:   link.Name,
			Source: link.Source,
		})
	}

	return pbLinks
}

func PbLink(link model.Link) *bookingpb.Link {
	return &bookingpb.Link{
		Id:     link.Id.String(),
		Name:   link.Name,
		Source: link.Source,
	}
}

func PbBookingList(bookings []uuid.UUID) []*bookingpb.Booking {
	pbBookings := make([]*bookingpb.Booking, 0, len(bookings))
	for _, bookingsUUID := range bookings {
		pbBookings = append(pbBookings, &bookingpb.Booking{
			Uuid: bookingsUUID.String(),
		})
	}

	return pbBookings
}

func toNotification(notification model.Notification) *bookingpb.Notification {
	return &bookingpb.Notification{
		Units:        notification.Units(),
		Fragments:    notification.Fragments(),
		StatusCode: bookingpb.Status_DEl,
	}
}
