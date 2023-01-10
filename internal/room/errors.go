package room

import "github.com/pkg/errors"

var (
	ErrRoomNotFound   = errors.New("room not found")
	ErrPermissionDenied  = errors.New("permission denied")
	ErrParticipantExists = errors.New("participant already exists")
)
