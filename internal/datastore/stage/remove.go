package stage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s Storage) Remove(ctx context.Context, stage model.Stage) error {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if err := s.deleteAttachments(ctx, tx, stage.Issues); err != nil {
			return err
		}

		if err := s.deleteStage(ctx, tx, stage.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if err := s.removeFromS3(ctx, stage.Issues); err != nil {
		return err
	}

	return nil
}

func (s Storage) deleteAttachments(ctx context.Context, tx pgx.Tx, issues []model.Issue) error {
	for _, issue := range issues {
		for _, attachment := range issue.Attachments {
			cmd, err := s.db.ExecTxBuilder(ctx, tx, attachmentsDeleteQuery(attachment.ID))
			if err != nil {
				return err
			}

			if cmd.RowsAffected() == 0 {
				return base.ErrNotFound
			}
		}
	}

	return nil
}

func (s Storage) deleteStage(ctx context.Context, tx pgx.Tx, stageID uuid.UUID) error {
	q, err := s.db.ExecTxBuilder(ctx, tx, deleteQuery(stageID))
	if err != nil {
		return err
	}

	if q.RowsAffected() == 0 {
		return base.ErrNotFound
	}

	return nil
}

func (s Storage) removeFromS3(ctx context.Context, issues []model.Issue) error {
	for _, issue := range issues {
		for _, attachment := range issue.Attachments {
			if err := s.fileStorage.RemoveFile(ctx, attachment.Key); err != nil {
				return err
			}
		}
	}

	return nil
}

func attachmentsDeleteQuery(attachmentID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(attachmentsTableName).Where("id = ?", attachmentID)
}
