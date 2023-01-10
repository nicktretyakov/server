package filters

import sq "github.com/Masterminds/squirrel"

type Filter interface {
	Apply(builder sq.SelectBuilder) sq.SelectBuilder
}
