package outmember

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/base"
	addressStore "be/internal/datastore/address"
	"be/internal/lib"
	"be/internal/model"
)

func (s storage) Create(
	ctx context.Context,
	agr model.Outmember,
	typeAddress addressStore.TypeAddress,
) (*model.Outmember, error) {
	agr.CreatedAt = s.db.Now()
	agr.ID = lib.UUID()

	if _, err := s.db.ExecBuilder(ctx, outmemberInsertQuery(agr, getNameTableQuery(typeAddress))); err != nil {
		return nil, err
	}

	return &agr, nil
}

func outmemberInsertQuery(agr model.Outmember, tableName string) sq.InsertBuilder {
	return base.Builder().
		Insert(tableName).
		Columns("id",
			"created_at",
			"type",
			"user_id",
			"address_id",
			"result",
			"extra",
			"role",
		).
		Values(
			agr.ID,
			agr.CreatedAt,
			agr.Type,
			agr.UserID,
			agr.AddressID,
			agr.Result,
			agr.Extra,
			agr.Role,
		)
}
