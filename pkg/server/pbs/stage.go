package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func Stage(stage model.Stage) *bookingpb.Stage {
	return &bookingpb.Stage{
		Uuid:     stage.ID.String(),
		Title:    stage.Title,
		Timeline: timelinePtr(stage.Timeline),
		Issues:   IssueList(stage.Issues),
	}
}

func StageList(stages []model.Stage) []*bookingpb.Stage {
	result := make([]*bookingpb.Stage, 0, len(stages))
	for _, stage := range stages {
		result = append(result, Stage(stage))
	}

	return result
}

func IssueList(issues []model.Issue) []*bookingpb.Issue {
	res := make([]*bookingpb.Issue, 0, len(issues))

	for _, i := range issues {
		res = append(res, Issue(i))
	}

	return res
}

func Issue(issue model.Issue) *bookingpb.Issue {
	return &bookingpb.Issue{
		Id:           issue.ID.String(),
		StageId:      issue.Stage.ID.String(),
		Title:        issue.Title,
		Description:  issue.Description,
		Timeline:     timelinePtr(issue.Timeline),
		Participants: PbUserList(issue.Participants),
		Attachments:  AttachmentList(issue.Attachments),
	}
}

func timelinePtr(timeline *model.Timeline) *bookingpb.Timeline {
	if timeline == nil {
		return nil
	}

	return &bookingpb.Timeline{
		Start: ToUTCString(timeline.StartAt),
		End:   ToUTCString(timeline.EndAt),
	}
}
