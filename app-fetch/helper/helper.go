package helper

import (
	"time"
)

func ConvertTimeStringToTime(timeString string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		layout := "2006-01-02T15:04:05-07:00"
		t, err := time.Parse(layout, timeString)
		if err != nil {
			layout := "2006-01-02T15:04:05-0700"
			t, _ := time.Parse(layout, timeString)

			return t
		}

		return t
	}

	return t
}
