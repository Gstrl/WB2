package usecase

import (
	"WB2/develop/dev11/internal/database"
	"WB2/develop/dev11/internal/models"
	"time"
)

type EventsForWeek struct {
	eventRep *database.CacheEventRep
}

func NewEventsForWeek(eventRep *database.CacheEventRep) *EventsForWeek {
	return &EventsForWeek{
		eventRep: eventRep,
	}
}

func (c *EventsForWeek) Execute(userID int) []models.Event {
	events := c.eventRep.GetEventsForWeek(userID, time.Now())
	return events
}
