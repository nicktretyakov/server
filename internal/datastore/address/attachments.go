package address

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) FindAttachmentByID(ctx context.Context, attachmentID uuid.UUID) (*model.Attachment, error) {
	var att dbmodel.Attachment

	if err := s.db.Get(ctx, attachmentsSelectQuery(attachmentID), &att); err != nil {
		return nil, err
	}

	return att.ToModelPtr(), nil
}

func (s *Storage) DeleteAttachment(ctx context.Context, attachment model.Attachment) error {
	return s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		cmd, err := s.db.ExecTxBuilder(ctx, tx, attachmentsDeleteQuery(attachment.ID))
		if err != nil {
			return err
		}

		if cmd.RowsAffected() == 0 {
			return base.ErrNotFound
		}

		return s.fileStorage.RemoveFile(ctx, attachment.Key)
	})
}

func (s *Storage) CreateAttachment(ctx context.Context, attachment model.Attachment) (*model.Attachment, error) {
	attachment.ID = lib.UUID()
	attachment.CreatedAt = s.db.Now()

	key, err := s.fileStorage.AddFile(ctx, attachment.Source(), attachment.Filename, attachment.Mime)
	if err != nil {
		return nil, err
	}

	attachment.Key = key

	if _, err := s.db.ExecBuilder(ctx, attachmentInsertQuery(attachment)); err != nil {
		return nil, err
	}

	uri, err := s.fileStorage.Link(attachment.Key, s.linkTimeLife)
	if err != nil {
		return nil, err
	}

	attachment.SetURL(uri)

	return &attachment, err
}

func (s *Storage) RenameAttachment(ctx context.Context, attachment model.Attachment) (*model.Attachment, error) {
	if err := s.fileStorage.RenameFile(ctx, attachment.Key, attachment.Filename, attachment.Mime); err != nil {
		return nil, err
	}

	if _, err := s.db.ExecBuilder(ctx, attachmentUpdateQuery(attachment)); err != nil {
		return nil, err
	}

	uri, err := s.fileStorage.Link(attachment.Key, s.linkTimeLife)
	if err != nil {
		return nil, err
	}

	attachment.SetURL(uri)

	return &attachment, err
}

func attachmentInsertQuery(attachment model.Attachment) sq.InsertBuilder {
	return base.Builder().
		Insert(attachmentsTableName).
		Columns(
			"id",
			"author_id",
			"created_at",
			"address_id",
			"key",
			"file_name",
			"mime",
			"size",
		).
		Values(
			attachment.ID,
			attachment.AuthorID,
			attachment.CreatedAt,
			attachment.AddressID,
			attachment.Key,
			attachment.Filename,
			attachment.Mime,
			attachment.Size,
		)
}

func attachmentUpdateQuery(attachment model.Attachment) sq.UpdateBuilder {
	return base.Builder().
		Update(attachmentsTableName).
		SetMap(
			map[string]interface{}{"file_name": attachment.Filename, "mime": attachment.Mime},
		).
		Where("id=?", attachment.ID)
}

func attachmentsSelectQuery(attachmentID uuid.UUID) sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"author_id",
			"created_at",
			"address_id",
			"key",
			"file_name",
			"mime",
			"size",
		).
		From(attachmentsTableName).
		Where("id=?", attachmentID)
}

func attachmentsDeleteQuery(attachmentID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(attachmentsTableName).Where("id = ?", attachmentID)
}
