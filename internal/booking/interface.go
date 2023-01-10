package booking

import (
	"context"

	"github.com/google/uuid"

	"be/internal/model"
)

type IBookingService interface {
	CreateBooking(ctx context.Context, user model.User, initialBooking model.Booking, portalCode uint64) (*model.Booking, error)
	UpdateBooking(ctx context.Context, user model.User, updatedBooking *model.Booking,
		portalCodeSupervisor, portalCodeAssignee uint64) (*model.Booking, error)
	SetBookingAssignee(ctx context.Context, book *model.Booking, portalCode uint64) (*model.Booking, error)
	SetBookingSupervisor(ctx context.Context, book *model.Booking, portalCode uint64) (*model.Booking, error)
	SendReport(ctx context.Context, user model.User, report model.ReportBooking) (*model.ReportBooking, error)
	SendFinalReport(ctx context.Context, user model.User, report model.FinalReport) (*model.FinalReport, error)
	AddBookingStageIssue(ctx context.Context, user model.User, stageID uuid.UUID) (*model.Issue, error)
	UpdateBookingStageIssue(ctx context.Context, user model.User, issueToUpdate model.Issue) (*model.Issue, error)
	RemoveBookingStageIssue(ctx context.Context, user model.User, issue model.Issue) error
}
