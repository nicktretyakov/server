package dictionary

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/model"
)

func (s storage) DepartmentList(ctx context.Context) ([]model.Department, error) {
	query := departmentSelectQuery()

	departments := make([]model.Department, 0)

	if err := s.db.Select(ctx, query, &departments); err != nil {
		return nil, err
	}

	return departments, nil
}

func departmentSelectQuery() sq.SelectBuilder {
	return base.Builder().
		Select("id", "title").
		From(DepartmentsTableName).
		OrderBy("title")
}
