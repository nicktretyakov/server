package stage

import (
	"github.com/google/uuid"
	bookingpb "be/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/model"
)

type issueRequest interface {
	GetTitle() string
	GetDescription() string
	GetTimeline() *bookingpb.Timeline
	GetParticipantPortalCodes() []uint32
	GetAttachmentId() []string
}

func validatedIssueRequest(req issueRequest) (issue model.Issue, err error) {
	timeline, err := getTimelineModel(req.GetTimeline())
	if err != nil {
		return issue, status.New(codes.InvalidArgument, "invalid date format").Err()
	}

	if !timeline.IsValid() {
		return issue, status.New(codes.InvalidArgument, "invalid timeline").Err()
	}

	users := make([]model.User, 0, len(req.GetParticipantPortalCodes()))

	for _, portalCode := range req.GetParticipantPortalCodes() {
		users = append(users, model.User{Employee: model.Employee{PortalCode: uint64(portalCode)}})
	}

	attachments := make([]model.Attachment, 0, len(req.GetAttachmentId()))

	for _, s := range req.GetAttachmentId() {
		attachmentID, parseErr := uuid.Parse(s)
		if parseErr != nil {
			return issue, parseErr
		}

		attachments = append(attachments, model.Attachment{ID: attachmentID})
	}

	issue.Participants = users
	issue.Title = req.GetTitle()
	issue.Description = req.GetDescription()
	issue.Timeline = timeline
	issue.Attachments = attachments

	return issue, nil
}
