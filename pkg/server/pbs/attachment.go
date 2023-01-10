package pbs

import (
	bookingpb "be/proto"

	"be/internal/model"
)

func Attachment(attachment model.Attachment) *bookingpb.Attachment {
	return &bookingpb.Attachment{
		Id:       attachment.ID.String(),
		Url:      attachment.URL(),
		FileName: attachment.Filename,
		Size:     uint64(attachment.Size),
	}
}

func AttachmentList(attachments []model.Attachment) []*bookingpb.Attachment {
	res := make([]*bookingpb.Attachment, 0, len(attachments))

	for _, attachment := range attachments {
		res = append(res, Attachment(attachment))
	}

	return res
}
