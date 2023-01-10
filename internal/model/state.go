package model

type State int

const (
	UnknownAddressState   State = 0
	PublishedAddressState State = 1 // Опубликовано
	ArchivedAddressState  State = 2 // Архивировано
)
