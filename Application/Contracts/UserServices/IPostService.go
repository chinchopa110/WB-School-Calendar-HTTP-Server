package UserServices

import (
	"WB2/Application/Domain"
)

type IPostService interface {
	AddUser(newUser *Domain.User) (*Domain.User, error)
	CreateEvent(userId int, date string, description string, key string) (Domain.Event, error)
	UpdateEventDate(userId int, eventId int, date string, key string) (Domain.Event, error)
	UpdateEventDescription(userId int, eventId int, description string, key string) (Domain.Event, error)
	DeleteEvent(userId int, eventId int, key string) (Domain.Event, error)
}
