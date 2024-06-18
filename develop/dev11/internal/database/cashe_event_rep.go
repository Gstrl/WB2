package database

import (
	"WB2/develop/dev11/internal/models"
	"errors"
	"sync"
	"time"
)

type CacheEventRep struct {
	cache         map[int]map[int]models.Event // Первый ключ userID, второй eventID
	autoIncrement map[int]int                  // Хранение автоинкремента для каждого userID
	mu            *sync.RWMutex
}

func NewCacheEventRep() *CacheEventRep {
	return &CacheEventRep{
		cache:         make(map[int]map[int]models.Event, 100),
		autoIncrement: make(map[int]int),
		mu:            new(sync.RWMutex),
	}
}

func (c *CacheEventRep) CreateEvent(event *models.Event) {
	defer c.mu.Unlock()
	c.mu.Lock()
	if _, ok := c.cache[event.UserID]; !ok {
		c.cache[event.UserID] = make(map[int]models.Event, 100)
	}
	event.EventID = c.autoIncrement[event.UserID]
	c.autoIncrement[event.UserID]++
	c.cache[event.UserID][event.EventID] = *event
}

func (c *CacheEventRep) getEventByID(userID, eventID int) (models.Event, error) {
	v, ok := c.cache[userID][eventID]
	if !ok {
		return models.Event{}, errors.New("error: can't find event")
	}
	return v, nil
}

func (c *CacheEventRep) UpdateEvent(userID, eventID int, event models.Event) error {
	defer c.mu.Unlock()
	c.mu.Lock()

	if _, err := c.getEventByID(userID, eventID); err != nil {
		return err
	}
	c.cache[userID][eventID] = event
	return nil
}

func (c *CacheEventRep) DeleteEvent(userID, eventID int) error {
	defer c.mu.Unlock()
	c.mu.Lock()
	if _, err := c.getEventByID(userID, eventID); err != nil {
		return err
	}
	delete(c.cache[userID], eventID)
	return nil
}

func (c *CacheEventRep) GetEventsForDay(userID int, date time.Time) []models.Event {
	defer c.mu.RUnlock()
	c.mu.RLock()

	eventsForDay := make([]models.Event, 0, len(c.cache[userID]))

	for _, v := range c.cache[userID] {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
			eventsForDay = append(eventsForDay, v)
		}
	}
	return eventsForDay
}

func (c *CacheEventRep) GetEventsForWeek(userID int, date time.Time) []models.Event {
	defer c.mu.RUnlock()
	c.mu.RLock()

	eventsForWeek := make([]models.Event, 0, len(c.cache[userID]))

	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	// Определяем пограничные дни недели для корректной работы функций time.After() и time.Before()
	startOfWeek := date.AddDate(0, 0, -weekday)
	endOfWeek := startOfWeek.AddDate(0, 0, 8)

	for _, v := range c.cache[userID] {
		if v.Date.After(startOfWeek) && v.Date.Before(endOfWeek) {
			eventsForWeek = append(eventsForWeek, v)
		}
	}

	return eventsForWeek
}

func (c *CacheEventRep) GetEventsForMonth(userID int, date time.Time) []models.Event {
	defer c.mu.RUnlock()
	c.mu.RLock()

	eventsForMonth := make([]models.Event, 0, 20)

	for _, v := range c.cache[userID] {
		if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
			eventsForMonth = append(eventsForMonth, v)
		}
	}

	return eventsForMonth
}
