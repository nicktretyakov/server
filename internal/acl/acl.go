package acl

import (
	"be/internal/model"
)

func IsHeadOfBooking(user model.User) bool {
	return user.Role == model.Admin
}

func CanCreateInitialOutmember(user model.User, io AddressACL) bool {
	return io.CanCreateInitialOutmember(user)
}

func CanCreateAcceptanceOutmember(user model.User, io AddressACL) bool {
	return io.CanCreateAcceptanceOutmember(user)
}

func CanAddAttachment(user model.User, io AddressACL) bool {
	return io.CanAddAttachment(user)
}

func CanRenameAttachment(user model.User, io AddressACL) bool {
	return io.CanAddAttachment(user)
}

func CanRemoveAttachment(user model.User, io AddressACL) bool {
	return io.CanRemoveAttachment(user)
}

func CanAddLink(user model.User, io AddressACL) bool {
	return io.CanAddLink(user)
}

func CanUpdateLink(user model.User, io AddressACL) bool {
	return io.CanUpdateLink(user)
}

func CanRemoveLink(user model.User, io AddressACL) bool {
	return io.CanRemoveLink(user)
}

func CanChangeStatus(user model.User, status model.Status, io AddressACL) bool {
	return io.CanChangeStatus(user, status)
}

func CanUpdateAddress(user model.User, io AddressACL) bool {
	return io.CanUpdateAddress(user)
}

func CanAddParticipants(user model.User, io AddressACL) bool {
	return io.CanAddParticipant(user)
}

func CanChangeState(user model.User, io AddressACL) bool {
	return io.CanChangeState(user)
}

func CanApproveFinalOutmember(user model.User, io AddressACL) bool {
	return io.CanApproveFinalOutmember(user)
}

func CanRegisterFinalOutmember(user model.User, io AddressACL) bool {
	return io.CanRegisterFinalOutmember(user)
}

func CanSendReport(user model.User, booking model.Booking, report model.ReportBooking) bool {
	return booking.CanSendReport(user, report)
}

func CanSendFinalReport(user model.User, booking model.Booking) bool {
	return booking.CanSendFinalReport(user)
}

func CanAddStage(user model.User, booking model.Booking) bool {
	return booking.CanAddStage(user)
}

func CanUpdateStage(user model.User, booking model.Booking) bool {
	return CanAddStage(user, booking)
}

func CanRemoveStage(user model.User, booking model.Booking) bool {
	return CanAddStage(user, booking)
}

func CanAddIssue(user model.User, booking model.Booking) bool {
	return CanAddStage(user, booking)
}

func IsUserWithoutPrivileges(user model.User, booking model.Booking) bool {
	return booking.IsUserWithoutPrivileges(user)
}
