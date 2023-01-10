package stage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	"be/internal/datastore/base"
	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s Storage) FindStageIssueByID(ctx context.Context, issueID uuid.UUID) (*model.Issue, error) {
	query := issueSelectQuery().Where("id=? and status>0", issueID)

	var rep dbmodel.Issue
	if err := s.db.Get(ctx, query, &rep); err != nil {
		return nil, err
	}

	attachments, err := s.issueAttachments(ctx, issueID)
	if err != nil {
		return nil, err
	}

	rep.Attachments = attachments

	return rep.ToModelPtr(), nil
}

func (s Storage) CreateIssue(ctx context.Context, issue model.Issue) (*model.Issue, error) {
	issue.ID = lib.UUID()
	issue.CreatedAt = s.db.Now()

	dbIssue := dbmodel.IssueFromModel(issue)

	if _, err := s.execIssue(ctx, dbIssue, issueInsertQuery(dbIssue)); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (s Storage) UpdateIssue(ctx context.Context, issue model.Issue) error {
	dbIssue := dbmodel.IssueFromModel(issue)

	cmd, err := s.execIssue(ctx, dbIssue, issueUpdateQuery(dbIssue))
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return base.ErrNotFound
	}

	return nil
}

func (s Storage) RemoveIssue(ctx context.Context, issue model.Issue) error {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if err := s.deleteAttachments(ctx, tx, []model.Issue{issue}); err != nil {
			return err
		}

		if err := s.deleteIssue(ctx, tx, issue.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if err := s.removeFromS3(ctx, []model.Issue{issue}); err != nil {
		return err
	}

	return nil
}

func (s Storage) deleteIssue(ctx context.Context, tx pgx.Tx, issueID uuid.UUID) error {
	cmd, err := s.db.ExecTxBuilder(ctx, tx, base.Builder().Delete(issueTableName).Where("id = ?", issueID))
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return base.ErrNotFound
	}

	return nil
}

func (s Storage) execIssue(ctx context.Context,
	issue dbmodel.Issue,
	builder sq.Sqlizer,
) (cmd pgconn.CommandTag, err error) {
	if err = s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		cmd, err = s.db.ExecTxBuilder(ctx, tx, builder)
		if err != nil {
			return err
		}

		if err = s.createIssueParticipants(ctx, tx, issue); err != nil {
			return err
		}

		return s.createIssueAttachments(ctx, tx, issue)
	}); err != nil {
		return cmd, err
	}

	return cmd, nil
}

func (s Storage) createIssueParticipants(ctx context.Context, tx pgx.Tx, issue dbmodel.Issue) error {
	deleteParticipantsQuery := base.Builder().Delete(issueParticipantsTableName).Where("stage_issue_id=?", issue.ID)
	if _, err := s.db.ExecTxBuilder(ctx, tx, deleteParticipantsQuery); err != nil {
		return err
	}

	if len(issue.Participants) == 0 {
		return nil
	}

	participantsInsertQuery := base.Builder().Insert(issueParticipantsTableName).
		Columns("stage_issue_id", "user_id")

	for _, participant := range issue.Participants {
		participantsInsertQuery = participantsInsertQuery.Values(issue.ID, participant.ID)
	}

	if _, err := s.db.ExecTxBuilder(ctx, tx, participantsInsertQuery); err != nil {
		return err
	}

	return nil
}

func (s *Storage) createIssueAttachments(ctx context.Context, tx pgx.Tx, issue dbmodel.Issue) error {
	deleteAttachmentsQuery := base.Builder().Delete(issueAttachmentsTableName).Where("stage_issue_id=?", issue.ID)
	if _, err := s.db.ExecTxBuilder(ctx, tx, deleteAttachmentsQuery); err != nil {
		return err
	}

	if len(issue.Attachments) == 0 {
		return nil
	}

	attachmentsInsertQuery := base.Builder().Insert(issueAttachmentsTableName).
		Columns("stage_issue_id", "booking_attachment_id")

	for _, attachment := range issue.Attachments {
		attachmentsInsertQuery = attachmentsInsertQuery.Values(issue.ID, attachment.ID)
	}

	if _, err := s.db.ExecTxBuilder(ctx, tx, attachmentsInsertQuery); err != nil {
		return err
	}

	return nil
}

func (s Storage) findStagesIssues(ctx context.Context, stagesID []uuid.UUID) ([]model.Issue, error) {
	res := make(dbmodel.IssueList, 0)
	query := issueSelectQuery().
		Where(sq.Eq{"stage_id": stagesID, "status": model.ActiveAggregateStatus}).
		OrderBy("start_at")

	if err := s.db.Select(ctx, query, &res); err != nil {
		return nil, err
	}

	issues := res.ToModels()

	for i, issue := range issues {
		attachments, err := s.issueAttachments(ctx, issue.ID)
		if err != nil {
			return nil, err
		}

		issues[i].Attachments = attachments.Attachments()
	}

	return issues, nil
}

func issueSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"id",
			"stage_id",
			"title",
			"description",
			"start_at",
			"end_at",
			"created_at",
			"participants",
			"status").
		LeftJoin(`(select sip.stage_issue_id, json_agg(row_to_json(u)) participants
              from stage_issue_participants sip join users u on sip.user_id = u.id
                group by stage_issue_id) x on x.stage_issue_id = stage_issues.id`).
		From(issueTableName)
}

func issueInsertQuery(issue dbmodel.Issue) sq.InsertBuilder {
	return base.Builder().
		Insert(issueTableName).
		Columns(
			"id",
			"stage_id",
			"title",
			"description",
			"start_at",
			"end_at",
			"created_at",
			"status").
		Values(
			issue.ID,
			issue.StageID,
			issue.Title,
			issue.Description,
			issue.StartAt,
			issue.EndAt,
			issue.CreatedAt,
			issue.Status,
		)
}

func issueUpdateQuery(issue dbmodel.Issue) sq.UpdateBuilder {
	return base.Builder().
		Update(issueTableName).
		SetMap(map[string]interface{}{
			"status":      issue.Status,
			"title":       issue.Title,
			"description": issue.Description,
			"start_at":    issue.StartAt,
			"end_at":      issue.EndAt,
		}).
		Where("id=?", issue.ID)
}

func (s *Storage) issueAttachments(ctx context.Context, issueID uuid.UUID) (dbmodel.AttachmentList, error) {
	var bookAttachments dbmodel.AttachmentList

	query := attachmentsSelectQuery().
		Join("stage_issue_attachments sia on attachments.id=sia.booking_attachment_id").
		Where("sia.stage_issue_id=?", issueID)

	if err := s.db.Select(ctx, query, &bookAttachments); err != nil {
		return nil, err
	}

	for i, att := range bookAttachments {
		uri, err := s.fileStorage.Link(att.Key, s.linkTimeLife)
		if err != nil {
			return nil, err
		}

		bookAttachments[i].URL = uri
	}

	return bookAttachments, nil
}

func attachmentsSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select(
			"attachments.id",
			"attachments.author_id",
			"attachments.created_at",
			"attachments.addressest_object_id",
			"attachments.key",
			"attachments.file_name",
			"attachments.mime",
			"attachments.size",
		).
		From(attachmentsTableName)
}
