package booking

import (
	"context"
	"unicode/utf8"

	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/model"
	"be/pkg/auth"
)

const (
	minTitleLen       = 3
	minDescriptionLen = 3
)

func (s Service) CreateInitialBooking(ctx context.Context,
	req *bookingpb.CreateInitialBookingRequest,
) (*bookingpb.CreateInitialBookingResponse, error) {
	user := auth.FromContext(ctx)

	newBooking, err := validateCreateDraftRequest(req)
	if err != nil {
		return nil, err
	}

	if req.PortalCodeSupervisor == 0 {
		return nil, status.New(codes.InvalidArgument, "empty supervisor portal code").Err()
	}

	createdBooking, err := s.bookingService.CreateBooking(ctx, *user, *newBooking, uint64(req.PortalCodeSupervisor))
	if err != nil {
		return nil, err
	}

	pbBooking, err := s.bookingFromDB(ctx, createdBooking.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.CreateInitialBookingResponse{
		Booking: pbBooking,
	}, nil
}

type iCreateInitialBookingRequest interface {
	GetType() bookingpb.BookingType
	GetTitle() string
	GetTimeline() *bookingpb.Timeline
	GetDescription() string
}

func validateCreateDraftRequest(request iCreateInitialBookingRequest) (*model.Booking, error) {
	p := model.Booking{}

	bookingType := model.BookingType(request.GetType())
	if !bookingType.IsValid() {
		return nil, status.New(codes.InvalidArgument, "invalid type").Err()
	}

	if utf8.RuneCountInString(request.GetTitle()) < minTitleLen {
		return nil, status.New(codes.InvalidArgument, "invalid title").Err()
	}

	if utf8.RuneCountInString(request.GetDescription()) < minDescriptionLen {
		return nil, status.New(codes.InvalidArgument, "invalid description").Err()
	}

	timeline, err := getTimelineModel(request.GetTimeline())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	if !timeline.IsValid() {
		return nil, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	p.Timeline = *timeline

	return &model.Booking{
		Type:        bookingType,
		Title:       request.GetTitle(),
		Timeline:    *timeline,
		Description: request.GetDescription(),
		Status:      model.InitialAddressStatus,
	}, nil
}
