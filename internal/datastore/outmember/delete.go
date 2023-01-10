package outmember

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/base"
	addressStore "be/internal/datastore/address"
	"be/internal/model"
)

func (s storage) BulkDelete(
	ctx context.Context,
	addressID uuid.UUID,
	agrType model.OutmemberType,
	typeAddress addressStore.TypeAddress,
) error {
	if _, err := s.db.ExecBuilder(ctx, deleteOutmemberBookingQuery(
		addressID,
		agrType,
		getNameTableQuery(typeAddress)),
	); err != nil {
		return err
	}

	return nil
}

func deleteOutmemberBookingQuery(addressID uuid.UUID, agrType model.OutmemberType, tableName string) sq.DeleteBuilder {
	return base.Builder().Delete(tableName).Where("address_id=? and type=?", addressID, agrType)
}
