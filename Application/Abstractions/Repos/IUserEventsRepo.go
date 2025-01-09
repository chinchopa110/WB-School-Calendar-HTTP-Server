package Repos

import (
	"WB2/Application/Domain"
)

type IUserEventsRepo interface {
	GetUserById(userId int) (*Domain.User, error)
	UpdateUser(User *Domain.User) error
	AddEvent(userId int, event *Domain.Event) error
	AddUser(User *Domain.User) error
	DeleteEvent(userId int, eventId int) error
}
