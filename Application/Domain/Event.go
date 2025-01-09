package Domain

type Event struct {
	Id          int
	Date        string
	Description string
}

func (e *Event) UpdateDate(newDate string) Event {
	e.Date = newDate
	return *e
}

func (e *Event) UpdateDescription(newDescription string) Event {
	e.Description = newDescription
	return *e
}
