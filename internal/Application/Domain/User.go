package Domain

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type User struct {
	Id     int     `json:"id"`
	Key    string  `json:"key"`
	Events []Event `json:"events"`
	mu     sync.Mutex
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(*u),
	})
}

func (u *User) IsKey(maybeKey string) bool {
	return u.Key == maybeKey
}

func (u *User) CreateEvent(event Event) {
	u.mu.Lock()
	u.Events = append(u.Events, event)
	u.mu.Unlock()
}

func (u *User) UpdateEventDate(id int, newDate string) (Event, error) {
	u.mu.Lock()
	for i, e := range u.Events {
		if e.Id == id {
			return u.Events[i].UpdateDate(newDate), nil
		}
	}
	u.mu.Unlock()
	return Event{}, errors.New("not found event")
}

func (u *User) UpdateEventDescription(id int, newDescription string) (Event, error) {
	u.mu.Lock()
	for i, e := range u.Events {
		if e.Id == id {
			return u.Events[i].UpdateDescription(newDescription), nil
		}
	}
	u.mu.Unlock()
	return Event{}, errors.New("not found event")
}

func (u *User) DeleteEvent(id int) (Event, error) {
	u.mu.Lock()

	for i, e := range u.Events {
		if e.Id == id {
			deletedEvent := u.Events[i]
			u.Events = append(u.Events[:i], u.Events[i+1:]...)
			return deletedEvent, nil
		}
	}
	u.mu.Unlock()
	return Event{}, errors.New("not found event")
}

func (u *User) EventsForDay() []Event {
	u.mu.Lock()
	var eventsForDay []Event
	today := time.Now().Format("2006-01-02")

	for _, e := range u.Events {
		if e.Date == today {
			eventsForDay = append(eventsForDay, e)
		}
	}
	u.mu.Unlock()
	return eventsForDay
}

func (u *User) EventsForWeek() []Event {
	u.mu.Lock()
	defer u.mu.Unlock()
	var eventsForWeek []Event
	start := time.Now().Truncate(24 * time.Hour)
	end := start.AddDate(0, 0, 7)

	for _, e := range u.Events {
		eventDate, _ := time.Parse("2006-01-02", e.Date)
		if !eventDate.Before(start) && eventDate.Before(end) {
			eventsForWeek = append(eventsForWeek, e)
		}
	}

	return eventsForWeek
}

func (u *User) EventsForMonth() []Event {
	u.mu.Lock()
	var eventsForMonth []Event
	now := time.Now()
	currentYear := now.Year()
	currentMonth := now.Month()

	for _, e := range u.Events {
		eventDate, _ := time.Parse("2006-01-02", e.Date)
		if eventDate.Year() == currentYear && eventDate.Month() == currentMonth {
			eventsForMonth = append(eventsForMonth, e)
		}
	}
	u.mu.Unlock()
	return eventsForMonth
}
