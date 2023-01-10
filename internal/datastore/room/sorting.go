package room

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

// PublishDateSorting Сортировка по дате публикации.
type PublishDateSorting struct {
	asc bool
}

func (s PublishDateSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("p.created_at %s", getSortType(s.asc)))
}

func NewSortingByPublishDate(asc bool) sorting.Sorting {
	return PublishDateSorting{asc: asc}
}

// CreationDateSorting Сортировка по фактической дате создания.
type CreationDateSorting struct {
	asc bool
}

func (s CreationDateSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("p.creation_date %s", getSortType(s.asc)))
}

func NewSortingByCreationDateSorting(asc bool) sorting.Sorting {
	return CreationDateSorting{asc: asc}
}
