package model

import (
	"time"
)

type Timeline struct {
	StartAt time.Time
	EndAt   time.Time
}

func (t Timeline) IsValid() bool {
	return t.EndAt.After(t.StartAt)
}

func (t Timeline) Periods() []Period {
	s := time.Date(t.StartAt.Year(), t.StartAt.Month(), 1, 0, 0, 0, 0, time.UTC)
	e := time.Date(t.EndAt.Year(), t.EndAt.Month(), 2, 0, 0, 0, 0, time.UTC)

	periods := make([]Period, 0)

	for e.After(s) {
		periods = append(periods, PeriodFromTime(s))
		s = s.AddDate(0, 1, 0)
	}

	return periods
}

func (t Timeline) ToUTC() Timeline {
	t.StartAt = t.StartAt.UTC()
	t.EndAt = t.EndAt.UTC()

	return t
}
