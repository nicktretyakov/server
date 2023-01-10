package model

import "github.com/google/uuid"

type NoteEvent string

const (
	UnknownNotify               NoteEvent = "unknown_notify"
	OnRegisterNotify            NoteEvent = "on_register_notify"
	OnAgreeNotify               NoteEvent = "on_agree_notify"
	ConfirmedNotify             NoteEvent = "confirmed_notify"
	DeclinedNotify              NoteEvent = "declined_notify"
	DoneNotify                  NoteEvent = "done_notify"
	NotSendReportNotify         NoteEvent = "not_send_report_notify"
	MissedReportNotify          NoteEvent = "missed_reports_notify"
	SentReportNotify            NoteEvent = "sent_report_notify"
	FinalReportOnRegisterNotify NoteEvent = "final_report_on_register_notify"
	FinalReportOnAgreeNotify    NoteEvent = "final_report_on_agree_notify"
	FinalReportDeclinedNotify   NoteEvent = "final_report_declined_notify"
)

type NoteStatus string

const (
	NotSend    NoteStatus = "not_send"
	Sent       NoteStatus = "sent"
	SendFailed NoteStatus = "send_failed"
)

type NoteObject string

const (
	AddressTypeBooking NoteObject = "booking"
	AddressTypeRoom NoteObject = "room"
)

type NoteSettings struct {
	UserID  *uuid.UUID
	EmailOn bool
	LifeOn  bool
}
