package booking

import (
	"context"

	"github.com/google/uuid"

	"be/internal/acl"
	"be/internal/model"
)

func (s Service) AddBookingStageIssue(ctx context.Context, user model.User, stageID uuid.UUID) (*model.Issue, error) {
	stage, err := s.store.Stage().FindByID(ctx, stageID)
	if err != nil {
		return nil, err
	}

	booking, err := s.store.Booking().FindByID(ctx, stage.BookingID)
	if err != nil {
		return nil, err
	}

	if !acl.CanAddIssue(user, *booking) {
		return nil, ErrPermissionDenied
	}

	issue := model.Issue{Stage: model.Stage{ID: stageID}, Status: model.InitialAggregateStatus}

	if err = s.findUsers(ctx, &issue); err != nil {
		return nil, err
	}

	return s.store.Stage().CreateIssue(ctx, issue)
}

func (s Service) UpdateBookingStageIssue(ctx context.Context, user model.User, issueToUpdate model.Issue) (*model.Issue, error) {
	issue, booking, err := s.findIssueBooking(ctx, issueToUpdate.ID)
	if err != nil {
		return nil, err
	}

	if !acl.CanAddIssue(user, *booking) {
		return nil, ErrPermissionDenied
	}

	if !checkAttachmentsIncluded(booking.Attachments, issueToUpdate.Attachments) {
		return nil, ErrAttachmentNotFound
	}

	issue.Title = issueToUpdate.Title
	issue.Description = issueToUpdate.Description
	issue.Timeline = issueToUpdate.Timeline
	issue.Participants = issueToUpdate.Participants
	issue.Attachments = issueToUpdate.Attachments
	issue.Status = model.ActiveAggregateStatus

	if err = s.findUsers(ctx, issue); err != nil {
		return nil, err
	}

	err = s.store.Stage().UpdateIssue(ctx, *issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func (s Service) RemoveBookingStageIssue(ctx context.Context, user model.User, issue model.Issue) error {
	_, booking, err := s.findIssueBooking(ctx, issue.ID)
	if err != nil {
		return err
	}

	if !acl.CanAddIssue(user, *booking) {
		return ErrPermissionDenied
	}

	return s.store.Stage().RemoveIssue(ctx, issue)
}

func (s Service) findIssueBooking(ctx context.Context, issueID uuid.UUID) (*model.Issue, *model.Booking, error) {
	issue, err := s.store.Stage().FindStageIssueByID(ctx, issueID)
	if err != nil {
		return nil, nil, err
	}

	stage, err := s.store.Stage().FindByID(ctx, issue.Stage.ID)
	if err != nil {
		return nil, nil, err
	}

	issue.Stage = *stage

	p, err := s.store.Booking().FindByID(ctx, stage.BookingID)

	return issue, p, err
}

func (s Service) findUsers(ctx context.Context, issue *model.Issue) (err error) {
	users := make([]model.User, 0, len(issue.Participants))

	for _, participant := range issue.Participants {
		u, err := s.userService.GetOrCreateUserByPortalCode(ctx, participant.Employee.PortalCode)
		if err != nil {
			return err
		}

		users = append(users, *u)
	}

	issue.Participants = users

	return nil
}

func checkAttachmentsIncluded(bookingAttachments []model.Attachment, attachments []model.Attachment) bool {
	bookingAttachmentsIDMap := make(map[uuid.UUID]bool)

	for _, attachment := range bookingAttachments {
		bookingAttachmentsIDMap[attachment.ID] = true
	}

	for _, attachment := range attachments {
		if _, ok := bookingAttachmentsIDMap[attachment.ID]; !ok {
			return false
		}
	}

	return true
}
