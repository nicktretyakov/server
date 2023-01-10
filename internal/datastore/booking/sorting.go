package booking

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

// Сортировка по дате публикации

type PublishDateSorting struct{ asc bool }

func (s PublishDateSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("created_at %s", getSortType(s.asc)))
}

func SortingByPublishDate(asc bool) sorting.Sorting {
	return PublishDateSorting{asc: asc}
}

// Сортировка по дате сдаче

type EndDateSorting struct{ asc bool }

func (s EndDateSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("timeline_end_at %s", getSortType(s.asc)))
}

func SortingByEndDate(asc bool) sorting.Sorting {
	return EndDateSorting{asc: asc}
}

// Сортировка по бюджету

type SlotSorting struct{ asc bool }

func (s SlotSorting) Apply(builder sq.SelectBuilder) sq.SelectBuilder {
	return builder.OrderBy(fmt.Sprintf("slot %s", getSortType(s.asc)))
}

func SortingBySlot(asc bool) sorting.Sorting {
	return SlotSorting{asc: asc}
}
