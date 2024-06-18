package timeparse

import "time"

func TimeParse(timeStr string) (time.Time, error) {
	const shortForm = "2006-01-02"
	t, err := time.Parse(shortForm, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
