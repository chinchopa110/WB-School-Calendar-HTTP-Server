package Domain

import "encoding/json"

type Event struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func (e *Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(*e),
	})
}

func (e *Event) UpdateDate(newDate string) Event {
	e.Date = newDate
	return *e
}

func (e *Event) UpdateDescription(newDescription string) Event {
	e.Description = newDescription
	return *e
}
