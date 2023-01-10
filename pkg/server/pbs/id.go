package pbs

import "github.com/google/uuid"

type IID interface {
	GetId() string
}

func ParseID(r IID) (uuid.UUID, error) {
	return uuid.Parse(r.GetId())
}
