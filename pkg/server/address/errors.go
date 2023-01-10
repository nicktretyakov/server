package address

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	addressService "be/internal/address"
	bookingService "be/internal/booking"
)

func addressErrorToStatus(err error) error {
	switch {
	case errIs(err, bookingService.ErrBookingNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errIs(err, addressService.ErrAttachmentNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errIs(err, addressService.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, err.Error())
	}

	return status.Error(codes.Unknown, err.Error())
}

func errIs(err, target error) bool {
	return errors.Is(err, target)
}
