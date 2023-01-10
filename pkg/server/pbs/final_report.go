package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func FinalReport(report model.FinalReport, attachments []model.Attachment) *bookingpb.FinalReport {
	return &bookingpb.FinalReport{
		Id:          report.ID.String(),
		Slot:      toNotification(report.Slot),
		EndDate:     ToUTCString(report.EndAt),
		Comment:     report.Comment,
		Status:      bookingpb.FinalReportStatus(report.Status),
		Attachments: AttachmentList(attachments),
	}
}
