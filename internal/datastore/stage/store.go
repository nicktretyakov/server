package stage

import (
	"time"

	"be/internal/datastore/base"
	"be/internal/filestorage"
)

const (
	stageTableName             = "booking_stages"
	stageTableNameAlias        = "booking_stages stages"
	issueTableName             = "stage_issues"
	issueParticipantsTableName = "stage_issue_participants"
	issueAttachmentsTableName  = "stage_issue_attachments"
	attachmentsTableName       = "attachments"
	linkTimeLife               = time.Hour
)

type Storage struct {
	db           *base.DB
	fileStorage  filestorage.IFileStorage
	linkTimeLife time.Duration
}

func New(db *base.DB, fileStorage filestorage.IFileStorage) *Storage {
	return &Storage{
		db: db, fileStorage: fileStorage,
		linkTimeLife: linkTimeLife,
	}
}
