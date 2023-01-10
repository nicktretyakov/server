package model

func (p Booking) CanCreateInitialOutmember(user User) bool {
	return p.IsHeadOfBooking(user) && p.Status == OnRegisterAddressStatus && p.State == PublishedAddressState
}

func (p Booking) CanCreateAcceptanceOutmember(user User) bool {
	isOnAgree := p.Status.Eq(OnAgreeAddressStatus)
	return isOnAgree && p.IsAssignee(user) && p.State == PublishedAddressState
}

func (p Booking) CanAddAttachment(user User) bool {
	return (p.canUpdateSupervisorOrAuthor(user) ||
		p.canUpdateHeadBooking(user) ||
		p.canUpdateAnyone(user, p.IsParticipant) ||
		p.canUpdateAnyone(user, p.IsAssignee) ||
		p.canUpdateAnyone(user, p.IsCEO) ||
		p.canUpdateAnyone(user, p.IsPartner)) && p.State == PublishedAddressState
}

func (p Booking) CanRenameAttachment(user User) bool {
	return p.CanAddAttachment(user)
}

func (p Booking) CanRemoveAttachment(user User) bool {
	return p.CanAddAttachment(user)
}

func (p Booking) CanAddLink(user User) bool {
	if p.IsFinalReport() {
		if p.FinalReport.Status.In(OnRegisterFinalReportStatus, OnAgreeFinalReportStatus) {
			return p.IsHeadOfBooking(user)
		}
	}

	if p.Status == InitialAddressStatus {
		return p.IsAuthor(user)
	}

	return (p.canUpdateSupervisorOrAuthor(user) ||
		p.canUpdateHeadBooking(user) ||
		p.canUpdateAnyone(user, p.IsParticipant) ||
		p.canUpdateAnyone(user, p.IsAssignee) ||
		p.canUpdateAnyone(user, p.IsCEO) ||
		p.canUpdateAnyone(user, p.IsPartner)) && p.State == PublishedAddressState
}

func (p Booking) CanUpdateLink(user User) bool {
	return p.CanAddLink(user)
}

func (p Booking) CanRemoveLink(user User) bool {
	return p.CanAddLink(user)
}

func (p Booking) CanChangeStatus(user User, status Status) bool {
	return (p.canUpdateSupervisorOrAuthor(user) ||
		p.canUpdateHeadBooking(user)) && status.Eq(OnRegisterAddressStatus)
}

func (p Booking) CanUpdateAddress(user User) bool {
	return (p.canUpdateHeadBooking(user) ||
		p.canUpdateBookingSupervisor(user) ||
		p.canUpdateBookingAuthor(user)) && p.State == PublishedAddressState
}

func (p Booking) CanCreateAssignee(user User) bool {
	return (p.canUpdateHeadBooking(user) ||
		p.canUpdateBookingSupervisor(user) ||
		p.canUpdateBookingAuthor(user)) && p.State == PublishedAddressState
}

func (p Booking) CanRegisterFinalOutmember(user User) bool {
	bookingFinalizeOnRegister := p.Status.Eq(FinalizeOnRegisterStatus)

	return p.IsHeadOfBooking(user) && bookingFinalizeOnRegister && p.State == PublishedAddressState
}

func (p Booking) CanApproveFinalOutmember(user User) bool {
	bookingFinalizeOnAgree := p.Status.Eq(FinalizeOnAgreeStatus)

	return p.IsAssignee(user) && bookingFinalizeOnAgree && p.State == PublishedAddressState
}

func (p Booking) CanSendReport(user User, report ReportBooking) bool {
	return (p.IsSupervisor(user) ||
		p.IsAuthor(user) ||
		p.IsHeadOfBooking(user) ||
		p.IsParticipant(user) ||
		p.IsAssignee(user)) && report.Status == NotSendReportStatus && p.State == PublishedAddressState
}

func (p Booking) CanAddParticipant(user User) bool {
	if p.FinalReport.Status.In(OnRegisterFinalReportStatus, OnAgreeFinalReportStatus) {
		return p.IsHeadOfBooking(user)
	}

	if p.Status == InitialAddressStatus {
		return p.IsAuthor(user)
	}

	return (p.IsHeadOfBooking(user) ||
		p.IsAuthor(user) ||
		p.IsSupervisor(user) ||
		p.canUpdateAnyone(user, p.IsParticipant) ||
		p.canUpdateAnyone(user, p.IsAssignee) ||
		p.canUpdateAnyone(user, p.IsCEO) ||
		p.canUpdateAnyone(user, p.IsPartner)) && p.State == PublishedAddressState
}

func (p Booking) CanSendFinalReport(user User) bool {
	userPerm := (p.IsSupervisor(user) || p.IsAuthor(user)) && (p.Status == ConfirmedAddressStatus || p.Status == FinalizeReportDeclined)

	return userPerm && p.State == PublishedAddressState
}

func (p Booking) IsUserWithoutPrivileges(user User) bool {
	return !p.IsHeadOfBooking(user) && !p.IsSupervisor(user) && !p.IsAuthor(user)
}

func (p Booking) CanAddStage(user User) bool {
	return (p.canUpdateSupervisorOrAuthor(user) ||
		p.canUpdateHeadBooking(user) ||
		p.canUpdateAnyone(user, p.IsParticipant) ||
		p.canUpdateAnyone(user, p.IsAssignee) ||
		p.canUpdateAnyone(user, p.IsCEO) ||
		p.canUpdateAnyone(user, p.IsPartner)) && p.State == PublishedAddressState
}

func (p Booking) CanChangeState(user User) bool {
	return p.IsHeadOfBooking(user)
}

func (p Booking) canUpdateSupervisorOrAuthor(user User) bool {
	supervisorPerm := p.IsSupervisor(user)
	authorPerm := p.IsAuthor(user)

	statusCheck := p.Status.In(
		InitialAddressStatus,
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		ConfirmedAddressStatus,
		OnAgreeAddressStatus,
		FinalizeReportDeclined,
	)

	return (supervisorPerm || authorPerm) && statusCheck
}

func (p Booking) canUpdateHeadBooking(user User) bool {
	headBookingPerm := p.IsHeadOfBooking(user)

	statusCheckHeadBooking := p.Status.In(
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		ConfirmedAddressStatus,
		OnAgreeAddressStatus,
		DoneAddressStatus,
		FinalizeOnRegisterStatus,
		FinalizeOnAgreeStatus,
		FinalizeReportDeclined,
	)

	return headBookingPerm && statusCheckHeadBooking
}

func (p Booking) canUpdateAnyone(user User, isWho func(u User) bool) bool {
	manPerm := isWho(user)
	statusCheck := p.Status.In(ConfirmedAddressStatus, FinalizeReportDeclined)

	return manPerm && statusCheck
}

func (p Booking) canUpdateBookingSupervisor(user User) bool {
	supervisorPerm := p.IsSupervisor(user)

	statusCheckSupervisor := p.Status.In(
		DeclinedAddressStatus,
		OnRegisterAddressStatus,
		OnAgreeAddressStatus,
		ConfirmedAddressStatus,
	)

	return supervisorPerm && statusCheckSupervisor
}

func (p Booking) canUpdateBookingAuthor(user User) bool {
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
