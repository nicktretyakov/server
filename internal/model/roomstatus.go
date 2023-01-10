package model

import "strconv"

type Status int

const (
	UnknownAddressStatus    Status = 0
	InitialAddressStatus    Status = 1 // Начальный
	DeclinedAddressStatus   Status = 2 // На доработку
	OnRegisterAddressStatus Status = 3 // На регистрации
	ConfirmedAddressStatus  Status = 4 // Согласовано
	DoneAddressStatus       Status = 5 // Завершен
	OnAgreeAddressStatus    Status = 6 // На согласовании
	FinalizeOnRegisterStatus     Status = 7 // Завершение регистрируется
	FinalizeOnAgreeStatus        Status = 8 // Завершение согласуется
	FinalizeReportDeclined       Status = 9 // Доработать завершение

	StatusUnknown            = "Неизвестный"
	StatusInitial            = "Начальный"
	StatusDeclined           = "На доработку"
	StatusOnRegister         = "На регистрации"
	StatusConfirmed          = "Согласован"
	StatusDone               = "Завершен"
	StatusOnAgree            = "На согласовании"
	StatusFinalizeOnRegister = "Завершение регистрируется"
	StatusFinalizeOnAgree    = "Завершение согласуется"
	StatusFinalizeDeclined   = "Доработать завершение"
)

func (s Status) Eq(s2 Status) bool {
	return s == s2
}

func (s Status) In(s2 ...Status) bool {
	for _, status := range s2 {
		if s.Eq(status) {
			return true
		}
	}

	return false
}

func (s Status) IsInitialOrDeclined() bool {
	return s.In(InitialAddressStatus, DeclinedAddressStatus)
}

func (s Status) String() string {
	return strconv.Itoa(int(s))
}

//nolint:gocyclo,cyclop
func (s Status) GetStatus() (output string) {
	switch s {
	case UnknownAddressStatus:
		output = StatusUnknown
	case InitialAddressStatus:
		output = StatusInitial
	case DeclinedAddressStatus:
		output = StatusDeclined
	case OnRegisterAddressStatus:
		output = StatusOnRegister
	case ConfirmedAddressStatus:
		output = StatusConfirmed
	case DoneAddressStatus:
		output = StatusDone
	case OnAgreeAddressStatus:
		output = StatusOnAgree
	case FinalizeOnRegisterStatus:
		output = StatusFinalizeOnRegister
	case FinalizeOnAgreeStatus:
		output = StatusFinalizeOnAgree
	case FinalizeReportDeclined:
		output = StatusFinalizeDeclined
	}

	return output
}
