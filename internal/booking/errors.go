package booking

import "github.com/pkg/errors"

var (
	ErrBookingNotFound    = errors.New("booking not found")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrAttachmentNotFound = errors.New("attachment not found")
)
