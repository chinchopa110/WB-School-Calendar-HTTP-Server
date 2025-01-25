package UserService

import (
	"WB2/Application/Abstractions/Repos"
	"WB2/Application/Domain"
	"errors"
)

type PostService struct {
	repo Repos.IUserEventsRepo
}

func CreatePostService(repo Repos.IUserEventsRepo) *PostService {
	return &PostService{repo: repo}
}

func (service PostService) AddUser(newUser *Domain.User) (*Domain.User, error) {
	var res = service.repo.AddUser(newUser)
	if res != nil {
		return nil, res
	}
	return newUser, nil
}

func (service PostService) CreateEvent(userId int, date string, description string, key string) (Domain.Event, error) {
	res, err := service.repo.GetUserById(userId)

	if err != nil {
		return Domain.Event{}, err
	}
	if !res.IsKey(key) {
		return Domain.Event{}, errors.New("INCORRECT KEY")
	}
	var event = Domain.Event{
		Date:        date,
		Description: description,
	}
	res.CreateEvent(event)

	var updateRes = service.repo.AddEvent(userId, &event)
	if updateRes != nil {
		return Domain.Event{}, updateRes
	}

	return event, nil
}

func (service PostService) UpdateEventDate(userId int, eventId int, date string, key string) (Domain.Event, error) {
	res, err := service.repo.GetUserById(userId)

	if err != nil {
		return Domain.Event{}, err
	}
	if !res.IsKey(key) {
		return Domain.Event{}, errors.New("INCORRECT KEY")
	}
	event, err := res.UpdateEventDate(eventId, date)
	if err != nil {
		return Domain.Event{}, err
	}
	var updateRes = service.repo.UpdateUser(res)
	if updateRes != nil {
		return Domain.Event{}, updateRes
	}

	return event, nil
}

func (service PostService) UpdateEventDescription(userId int, eventId int, description string, key string) (Domain.Event, error) {
	res, err := service.repo.GetUserById(userId)

	if err != nil {
		return Domain.Event{}, err
	}
	if !res.IsKey(key) {
		return Domain.Event{}, errors.New("INCORRECT KEY")
	}
	event, err := res.UpdateEventDescription(eventId, description)
	if err != nil {
		return Domain.Event{}, err
	}
	var updateRes = service.repo.UpdateUser(res)
	if updateRes != nil {
		return Domain.Event{}, updateRes
	}

	return event, nil
}

func (service PostService) DeleteEvent(userId int, eventId int, key string) (Domain.Event, error) {
	res, err := service.repo.GetUserById(userId)

	if err != nil {
		return Domain.Event{}, err
	}
	if !res.IsKey(key) {
		return Domain.Event{}, errors.New("INCORRECT KEY")
	}
	event, err := res.DeleteEvent(eventId)
	if err != nil {
		return Domain.Event{}, err
	}

	var updateRes = service.repo.DeleteEvent(userId, eventId)
	if updateRes != nil {
		return Domain.Event{}, updateRes
	}
	return event, nil
}
