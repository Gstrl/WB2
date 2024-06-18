package models

import "time"

type Event struct {
	UserID  int
	EventID int
	Title   string
	Date    time.Time
}

func NewEvent(userID int, title string, date time.Time) *Event {
	return &Event{
		UserID: userID,
		Title:  title,
		Date:   date,
	}
}
