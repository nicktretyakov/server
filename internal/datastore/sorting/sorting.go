package sorting

import sq "github.com/Masterminds/squirrel"

type Sorting interface {
	Apply(builder sq.SelectBuilder) sq.SelectBuilder
}
