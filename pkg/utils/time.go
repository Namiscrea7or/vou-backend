package utils

import "time"

func Now() time.Time {
	return time.Now().Truncate(time.Second).UTC()
}
