package usecase

import (
	"WB2/develop/dev11/internal/database"
)

type DeleteEvent struct {
	eventRep *database.CacheEventRep
}

func NewDeleteEvent(eventRep *database.CacheEventRep) *DeleteEvent {
	return &DeleteEvent{
		eventRep: eventRep,
	}
}

func (c *DeleteEvent) Execute(userID, eventID int) error {
	err := c.eventRep.DeleteEvent(userID, eventID)
	if err != nil {
		return err
	}
	return nil
}
