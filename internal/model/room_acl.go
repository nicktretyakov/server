package model

func (p Room) IsHeadOfBooking(user User) bool {
	return user.Role == Admin
}

func (p Room) CanCreateInitialOutmember(user User) bool {
	return p.IsHeadOfBooking(user) && p.Status == OnRegisterAddressStatus && p.State == PublishedAddressState
}

func (p Room) CanCreateAcceptanceOutmember(user User) bool {
	isOnAgree := p.Status.Eq(OnAgreeAddressStatus)
	return isOnAgree && p.IsAssignee(user) && p.State == PublishedAddressState
}

func (p Room) CanRegisterFinalOutmember(_ User) bool {
	return false
}

func (p Room) CanApproveFinalOutmember(_ User) bool {
	return false
}

func (p Room) CanAddAttachment(user User) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanRenameAttachment(user User) bool {
	return p.CanAddAttachment(user)
}

func (p Room) CanRemoveAttachment(user User) bool {
	return p.CanAddAttachment(user)
}

func (p Room) CanAddLink(user User) bool {
	if p.Status == InitialAddressStatus {
		return p.IsAuthor(user)
	}

	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanUpdateLink(user User) bool {
	return p.CanAddLink(user)
}

func (p Room) CanRemoveLink(user User) bool {
	return p.CanAddLink(user)
}

func (p Room) CanChangeStatus(user User, status Status) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.canUpdateRoomParticipant(user) ||
		p.IsHeadOfBooking(user)) && status.Eq(OnRegisterAddressStatus) && p.State == PublishedAddressState
}

func (p Room) CanUpdateAddress(user User) bool {
	return (p.canUpdateHeadBookingRoom(user) ||
		p.canUpdateRoomOwner(user) ||
		p.canUpdateRoomAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanAddParticipant(user User) bool {
	if p.Status == InitialAddressStatus {
		return p.IsAuthor(user)
	}

	return (p.IsHeadOfBooking(user) ||
		p.IsAuthor(user) ||
		p.IsOwner(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanChangeState(user User) bool {
	return p.IsHeadOfBooking(user)
}

func (p Room) CanUpdateSlot(user User) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanUpdateEquipment(user User) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanUpdateRelease(user User) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) CanSendReport(user User) bool {
	return (p.canUpdateRoomOwnerOrAuthor(user) ||
		p.canUpdateRoomAssignee(user) ||
		p.IsHeadOfBooking(user) ||
		p.canUpdateRoomParticipant(user)) && p.State == PublishedAddressState
}

func (p Room) canUpdateRoomOwnerOrAuthor(user User) bool {
	authorPerm := p.IsAuthor(user)
	ownerPerm := p.IsOwner(user)

	statusCheck := p.Status.In(
		InitialAddressStatus,
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		ConfirmedAddressStatus,
		OnAgreeAddressStatus,
	)

	return (ownerPerm || authorPerm) && statusCheck
}

func (p Room) canUpdateRoomAssignee(user User) bool {
	isConfirmed := p.Status.Eq(ConfirmedAddressStatus)
	return isConfirmed && p.IsAssignee(user)
}

func (p Room) canUpdateRoomParticipant(user User) bool {
	isConfirmed := p.Status.Eq(ConfirmedAddressStatus)
	return isConfirmed && p.IsParticipant(user)
}

func (p Room) canUpdateHeadBookingRoom(user User) bool {
	headBookingPerm := p.IsHeadOfBooking(user)

	statusCheckHeadBooking := p.Status.In(
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		ConfirmedAddressStatus,
		OnAgreeAddressStatus,
		DoneAddressStatus,
	)

	return headBookingPerm && statusCheckHeadBooking
}

func (p Room) canUpdateRoomOwner(user User) bool {
	ownerPerm := p.IsOwner(user)

	statusCheckSupervisor := p.Status.In(
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		OnAgreeAddressStatus,
		ConfirmedAddressStatus,
	)

	return ownerPerm && statusCheckSupervisor
}

func (p Room) canUpdateRoomAuthor(user User) bool {
	authorPerm := p.IsAuthor(user)

	statusCheckAuthor := p.Status.In(
		InitialAddressStatus,
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		OnAgreeAddressStatus,
		ConfirmedAddressStatus,
	)

	return authorPerm && statusCheckAuthor
}
