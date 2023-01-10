package validators

import "time"

func Time(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
