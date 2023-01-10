package pbs

import "time"

const TimeLayout = time.RFC3339

func ToUTCString(t time.Time) string {
	return t.In(time.UTC).Format(TimeLayout)
}
