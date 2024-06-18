package usecase

import (
	"WB2/develop/dev11/internal/database"
	"WB2/develop/dev11/internal/models"
)

type UpdateEvent struct {
	eventRep *database.CacheEventRep
}

func NewUpdateEvent(eventRep *database.CacheEventRep) *UpdateEvent {
	return &UpdateEvent{
		eventRep: eventRep,
	}
}

func (c *UpdateEvent) Execute(userID, eventID int, event models.Event) error {
	err := c.eventRep.UpdateEvent(userID, eventID, event)
	if err != nil {
		return err
	}
	return nil
}
