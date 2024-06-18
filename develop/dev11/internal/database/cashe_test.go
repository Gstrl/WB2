package database

import (
	"WB2/develop/dev11/internal/models"
	"sync"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	rep := NewCacheEventRep()
	event := models.NewEvent(1, "Test Event", time.Now())
	rep.CreateEvent(event)

	if len(rep.cache[1]) != 1 {
		t.Errorf("Expected 1 event, got %d", len(rep.cache[1]))
	}

	if event.EventID != 0 {
		t.Errorf("Expected EventID to be 0, got %d", event.EventID)
	}
}

func TestGetEventByID(t *testing.T) {
	rep := NewCacheEventRep()
	event := models.NewEvent(1, "Test Event", time.Now())
	rep.CreateEvent(event)

	retrievedEvent, err := rep.getEventByID(1, event.EventID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if retrievedEvent.Title != event.Title {
		t.Errorf("Expected title %s, got %s", event.Title, retrievedEvent.Title)
	}
}

func TestUpdateEvent(t *testing.T) {
	rep := NewCacheEventRep()
	event := models.NewEvent(1, "Test Event", time.Now())
	rep.CreateEvent(event)

	updatedEvent := models.NewEvent(1, "Updated Event", time.Now())

	err := rep.UpdateEvent(1, event.EventID, *updatedEvent)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	retrievedEvent, err := rep.getEventByID(1, event.EventID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if retrievedEvent.Title != updatedEvent.Title {
		t.Errorf("Expected title %s, got %s", updatedEvent.Title, retrievedEvent.Title)
	}
}

func TestDeleteEvent(t *testing.T) {
	rep := NewCacheEventRep()
	event := models.NewEvent(1, "Test Event", time.Now())
	rep.CreateEvent(event)

	err := rep.DeleteEvent(1, event.EventID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = rep.getEventByID(1, event.EventID)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetEventsForDay(t *testing.T) {
	rep := NewCacheEventRep()
	date := time.Now()
	event := models.NewEvent(1, "Test Event", date)
	rep.CreateEvent(event)

	events := rep.GetEventsForDay(1, date)
	if len(events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(events))
	}

	if events[0].Title != event.Title {
		t.Errorf("Expected title %s, got %s", event.Title, events[0].Title)
	}
}

func TestGetEventsForWeek(t *testing.T) {
	rep := NewCacheEventRep()
	date := time.Now()
	event := models.NewEvent(1, "Test Event", date)
	rep.CreateEvent(event)

	events := rep.GetEventsForWeek(1, date)
	if len(events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(events))
	}

	if events[0].Title != event.Title {
		t.Errorf("Expected title %s, got %s", event.Title, events[0].Title)
	}
}

func TestGetEventsForMonth(t *testing.T) {
	rep := NewCacheEventRep()
	date := time.Now()
	event := models.NewEvent(1, "Test Event", date)
	rep.CreateEvent(event)

	events := rep.GetEventsForMonth(1, date)
	if len(events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(events))
	}

	if events[0].Title != event.Title {
		t.Errorf("Expected title %s, got %s", event.Title, events[0].Title)
	}
}

func TestConcurrentAccess(t *testing.T) {
	rep := NewCacheEventRep()
	event := models.NewEvent(1, "Concurrent Event", time.Now())

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rep.CreateEvent(event)
		}()
	}
	wg.Wait()

	if len(rep.cache[1]) != 100 {
		t.Errorf("Expected 100 events, got %d", len(rep.cache[1]))
	}
}
