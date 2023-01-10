package model

import "time"

type Period struct {
	Year  int32
	Month int32
}

func (p Period) Time() time.Time {
	return time.Date(int(p.Year), time.Month(p.Month), 1, 0, 0, 0, 0, time.UTC)
}

func PeriodFromTime(t time.Time) Period {
	return Period{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
	}
}
