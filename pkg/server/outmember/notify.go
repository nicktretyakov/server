package outmember

import "be/internal/model"

//nolint:exhaustive
func getNoteTypeByStatus(toStatus model.Status) model.NoteEvent {
	switch toStatus {
	case model.DeclinedAddressStatus: // На доработку
		return model.DeclinedNotify
	case model.OnRegisterAddressStatus: // На регистрации
		return model.OnRegisterNotify
	case model.ConfirmedAddressStatus: // Согласовано
		return model.ConfirmedNotify
	case model.DoneAddressStatus: // Завершен
		return model.DoneNotify
	case model.OnAgreeAddressStatus: // На согласовании
		return model.OnAgreeNotify
	default:
		return model.UnknownNotify
	}
}
