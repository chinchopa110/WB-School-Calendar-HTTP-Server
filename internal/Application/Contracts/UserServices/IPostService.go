package UserServices

import (
	Domain2 "WB2/internal/Application/Domain"
)

type IPostService interface {
	AddUser(newUser *Domain2.User) (*Domain2.User, error)
	CreateEvent(userId int, date string, description string, key string) (Domain2.Event, error)
	UpdateEventDate(userId int, eventId int, date string, key string) (Domain2.Event, error)
	UpdateEventDescription(userId int, eventId int, description string, key string) (Domain2.Event, error)
	DeleteEvent(userId int, eventId int, key string) (Domain2.Event, error)
}
