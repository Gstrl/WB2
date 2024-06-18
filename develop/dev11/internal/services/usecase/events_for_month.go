package usecase

import (
	"WB2/develop/dev11/internal/database"
	"WB2/develop/dev11/internal/models"
	"time"
)

type EventsForMonth struct {
	eventRep *database.CacheEventRep
}

func NewEventsForMonth(eventRep *database.CacheEventRep) *EventsForMonth {
	return &EventsForMonth{
		eventRep: eventRep,
	}
}

func (c *EventsForMonth) Execute(userID int) []models.Event {
	events := c.eventRep.GetEventsForMonth(userID, time.Now())
	return events
}
