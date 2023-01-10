package stage

import (
	"context"

	"be/internal/datastore/dbmodel"
	"be/internal/lib"
	"be/internal/model"
)

func (s *Storage) Create(ctx context.Context, item model.Stage) (*model.Stage, error) {
	dbStage := dbmodel.StageFromModel(item)

	dbStage.ID = lib.UUID()
	dbStage.CreatedAt = s.db.Now()

	if _, err := s.db.ExecBuilder(ctx, insertQuery(dbStage)); err != nil {
		return nil, err
	}

	return dbStage.ToModelPtr(), nil
}
