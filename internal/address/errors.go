package address

import "github.com/pkg/errors"

var (
	ErrAttachmentNotFound      = errors.New("attachment not found")
	ErrPermissionDenied        = errors.New("permission denied")
	ErrUnknownTypeAddress = errors.New("unknown invest object type")
)
