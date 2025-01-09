package UserServices

import (
	"WB2/Application/Domain"
)

type IGetService interface {
	EventsForDay(userId int, key string) ([]Domain.Event, error)
	EventsForWeek(userId int, key string) ([]Domain.Event, error)
	EventsForMonth(userId int, key string) ([]Domain.Event, error)
}
