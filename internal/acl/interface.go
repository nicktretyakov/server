package acl

import "be/internal/model"

type AddressACL interface {
	CanCreateInitialOutmember(user model.User) bool
	CanCreateAcceptanceOutmember(user model.User) bool
	CanAddAttachment(user model.User) bool
	CanRenameAttachment(user model.User) bool
	CanRemoveAttachment(user model.User) bool
	CanAddLink(user model.User) bool
	CanUpdateLink(user model.User) bool
	CanRemoveLink(user model.User) bool
	CanChangeStatus(user model.User, status model.Status) bool
	CanUpdateAddress(user model.User) bool
	CanRegisterFinalOutmember(user model.User) bool
	CanApproveFinalOutmember(user model.User) bool
	CanAddParticipant(user model.User) bool
	CanChangeState(user model.User) bool
}
