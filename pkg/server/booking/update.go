package booking

import (
	"context"

	bookingpb "be/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/lib"
	"be/internal/model"
	"be/pkg/auth"
	"be/pkg/server/validators"
)

func (s Service) UpdateBooking(ctx context.Context, request *bookingpb.UpdateBookingRequest) (*bookingpb.UpdateBookingResponse, error) { //nolint:lll
	updReq, err := validateUpdateRequest(request)
	if err != nil {
		return nil, err
	}

	bookingToUpdate, err := s.store.Booking().FindByID(ctx, updReq.BookingID)
	if err != nil {
		return nil, err
	}

	user := auth.FromContext(ctx)

	updReq.update(bookingToUpdate)

	if _, err = s.bookingService.UpdateBooking(
		ctx,
		*user,
		bookingToUpdate,
		uint64(request.GetPortalCodeSupervisor()),
		uint64(request.GetPortalCodeAssignee()),
	); err != nil {
		return nil, err
	}

	pbBooking, err := s.bookingFromDB(ctx, bookingToUpdate.ID)
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateBookingResponse{
		Booking: pbBooking,
	}, nil
}

func validateUpdateRequest(req *bookingpb.UpdateBookingRequest) (*updateRequest, error) {
	bookingID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid booking_id").Err()
	}

	if len(req.GetDepartments()) == 0 {
		return nil, status.New(codes.InvalidArgument, "departments empty").Err()
	}

	departmentIDs, err := lib.ParseUUIDStrings(req.GetDepartments())
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "invalid department_id").Err()
	}

	departments := make([]model.Department, 0, len(departmentIDs))
	for _, depID := range departmentIDs {
		departments = append(departments, model.Department{ID: depID})
	}

	slot, err := validators.Notification(req.GetSlot())
	if err != nil {
		return nil, status.Newf(codes.InvalidArgument, "slot invalid: %s", err.Error()).Err()
	}

	book, err := validateCreateDraftRequest(req)
	if err != nil {
		return nil, err
	}

	return &updateRequest{
		Type:        book.Type,
		Title:       book.Title,
		City:        req.GetCity(),
		Timeline:    book.Timeline,
		Description: book.Description,
		BookingID:   bookingID,
		Slot:        slot,
		Goal:        req.GetGoal(),
		Departments: departments,
	}, nil
}

// Update values in a request.
type updateRequest struct {
	BookingID   uuid.UUID
	Slot        model.Notification
	Goal        string
	Departments []model.Department
	Type        model.BookingType
	Title       string
	City        string
	Timeline    model.Timeline
	Description string
}

func (r updateRequest) update(book *model.Booking) {
	book.Slot = r.Slot
	book.Goal = r.Goal
	book.Departments = r.Departments
	book.Type = r.Type
	book.Title = r.Title
	book.City = r.City
	book.Timeline = r.Timeline
	book.Description = r.Description
}
