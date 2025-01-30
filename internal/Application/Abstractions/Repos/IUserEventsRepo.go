package Repos

import (
	Domain2 "WB2/internal/Application/Domain"
)

type IUserEventsRepo interface {
	GetUserById(userId int) (*Domain2.User, error)
	UpdateUser(User *Domain2.User) error
	AddEvent(userId int, event *Domain2.Event) error
	AddUser(User *Domain2.User) error
	DeleteEvent(userId int, eventId int) error
}
