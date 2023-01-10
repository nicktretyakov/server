package outmember

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	"be/internal/datastore/filters"
	addressStore "be/internal/datastore/address"
	"be/internal/model"
)

func (s storage) List(
	ctx context.Context,
	typeAddress addressStore.TypeAddress,
	filters ...filters.Filter,
) ([]model.Outmember, error) {
	query := selectQueryOutmember(getNameTableQuery(typeAddress))
	for _, f := range filters {
		query = f.Apply(query)
	}

	var outmembersList []model.Outmember

	if err := s.db.Select(ctx, query, &outmembersList); err != nil {
		return nil, err
	}

	return outmembersList, nil
}

func selectQueryOutmember(nameTable string) sq.SelectBuilder {
	return base.Builder().
		Select("id",
			"created_at",
			"type",
			"user_id",
			"address_id",
			"result",
			"extra",
			"role",
		).
		From(nameTable).
		OrderBy("created_at desc")
}
