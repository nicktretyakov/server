package room

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

func (s *Storage) AddReleases(ctx context.Context, releases []model.Release) ([]model.Release, error) {
	if err := s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		for i := range releases {
			releases[i].ID = lib.UUID()
			releaseDBModel := dbmodel.ReleaseFromModel(releases[i])
			if _, err := s.db.ExecTxBuilder(ctx, tx, releaseInsertQuery(releaseDBModel)); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return releases, nil
}

func (s *Storage) UpdateRelease(ctx context.Context, release model.Release) (*model.Release, error) {
	releaseForDB := dbmodel.ReleaseFromModel(release)
	if _, err := s.db.ExecBuilder(ctx, releaseUpdateQuery(releaseForDB)); err != nil {
		return nil, err
	}

	return &release, nil
}

func (s *Storage) DeleteRelease(ctx context.Context, releaseID uuid.UUID) error {
	return s.db.Conn.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		cmd, err := s.db.ExecTxBuilder(ctx, tx, releaseDeleteQuery(releaseID))
		if err != nil {
			return err
		}

		if cmd.RowsAffected() == 0 {
			return base.ErrNotFound
		}

		return nil
	})
}

func releaseInsertQuery(release dbmodel.Release) sq.InsertBuilder {
	return base.Builder().Insert(releaseTableName).
		Columns(
			"id",
			"room_id",
			"title",
			"description",
			"date",
			"fact_slot",
		).
		Values(
			release.ID,
			release.RoomID,
			release.Title,
			release.Description,
			release.Date,
			release.FactSlot,
		)
}

func releaseUpdateQuery(release dbmodel.Release) sq.UpdateBuilder {
	return base.Builder().
		Update(releaseTableName).
		SetMap(map[string]interface{}{
			"title":       release.Title,
			"description": release.Description,
			"date":        release.Date,
			"fact_slot": release.FactSlot,
		}).
		Where("id = ?", release.ID)
}

func releaseDeleteQuery(releaseID uuid.UUID) sq.DeleteBuilder {
	return base.Builder().Delete(releaseTableName).Where("id = ?", releaseID)
}
