package address

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"be/internal/datastore/sorting"
)

func getSortType(asc bool) string {
	if asc {
		return "asc"
	}

	return "desc"
}

type CreatedAtSorting struct {
	asc bool
}

func (s CreatedAtSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("created_at %s", getSortType(s.asc)))
}

func NewSortingByCreatedAt(asc bool) sorting.Sorting {
	return CreatedAtSorting{asc: asc}
}

type ReadAtSorting struct {
	asc bool
}

func (s ReadAtSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("read_at %s", getSortType(s.asc)))
}

func NewSortingByReadAt(asc bool) sorting.Sorting {
	return ReadAtSorting{asc: asc}
}
