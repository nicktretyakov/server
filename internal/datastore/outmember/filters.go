package outmember

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"be/internal/datastore/filters"
	"be/internal/model"
)

type filter struct {
	apply func(builder sq.SelectBuilder) sq.SelectBuilder
}

func (f filter) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return f.apply(builder)
}

func NewAddressIDFilter(addressID uuid.UUID) filters.Filter {
	return filter{func(builder sq.SelectBuilder) sq.SelectBuilder {
		return builder.Where("address_id=?", addressID)
	}}
}

func NewOutmemberTypeFilter(outmemberType model.OutmemberType) filters.Filter {
	return filter{func(builder sq.SelectBuilder) sq.SelectBuilder {
		return builder.Where("type=?", outmemberType)
	}}
}
