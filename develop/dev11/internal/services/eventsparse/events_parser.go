package eventsparse

import (
	"WB2/develop/dev11/internal/models"
	"fmt"
	"strings"
)

func EventsPars(events []models.Event) string {
	var strEvents strings.Builder
	for _, event := range events {
		strEvent := fmt.Sprintf("title - %s event_id - %d : Date - %s, ", event.Title, event.EventID, event.Date.String())
		strEvents.WriteString(strEvent)
	}
	return strEvents.String()
}
