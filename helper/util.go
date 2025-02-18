package helper

import "time"

func StringISOToDateTime(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}
