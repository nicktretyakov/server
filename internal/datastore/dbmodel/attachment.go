package dbmodel

import (
	"time"

	"github.com/google/uuid"

	"be/internal/model"
)

type Attachment struct {
	ID             uuid.UUID `db:"id" yaml:"id"`
	AuthorID       uuid.UUID `db:"author_id" yaml:"author_id"`
	CreatedAt      time.Time `db:"created_at" yaml:"created_at"`
	AddressID uuid.UUID `db:"address_id" yaml:"address_id"`
	Key            string    `db:"key" yaml:"key"`
	Filename       string    `db:"file_name" yaml:"file_name"`
	Mime           string    `db:"mime" yaml:"mime"`
	Size           int       `db:"size" yaml:"size"`
	URL            string    `db:"-"`
}

func (a Attachment) ToModel() model.Attachment {
	att := model.Attachment{
		ID:             a.ID,
		AuthorID:       a.AuthorID,
		CreatedAt:      a.CreatedAt,
		AddressID: a.AddressID,
		Key:            a.Key,
		Filename:       a.Filename,
		Mime:           a.Mime,
		Size:           a.Size,
	}

	att.SetURL(a.URL)

	return att
}

func (a Attachment) ToModelPtr() *model.Attachment {
	aPtr := a.ToModel()

	return &aPtr
}

func AttachmentFromModel(a model.Attachment) Attachment {
	return Attachment{
		ID:             a.ID,
		AuthorID:       a.AuthorID,
		CreatedAt:      a.CreatedAt,
		AddressID: a.AddressID,
		Key:            a.Key,
		Filename:       a.Filename,
		Mime:           a.Mime,
		Size:           a.Size,
		URL:            a.URL(),
	}
}

type AttachmentList []Attachment

func (l AttachmentList) Attachments() []model.Attachment {
	modelsList := make([]model.Attachment, 0, len(l))
	for _, att := range l {
		modelsList = append(modelsList, att.ToModel())
	}

	return modelsList
}

func AttachmentListFromModel(attachments []model.Attachment) AttachmentList {
	res := make(AttachmentList, 0, len(attachments))

	for _, attachment := range attachments {
		res = append(res, AttachmentFromModel(attachment))
	}

	return res
}
