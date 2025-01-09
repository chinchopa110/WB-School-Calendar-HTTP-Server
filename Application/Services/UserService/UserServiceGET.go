package UserService

import (
	"WB2/Application/Abstractions/Repos"
	"WB2/Application/Domain"
	"errors"
)

type GetService struct {
	repo Repos.IUserEventsRepo
}

func CreateGetService(repo Repos.IUserEventsRepo) *GetService {
	return &GetService{repo: repo}
}

func (service *GetService) EventsForDay(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	if len(user.EventsForDay()) == 0 {
		return nil, errors.New("Events not found ")
	}

	return user.EventsForDay(), nil
}

func (service *GetService) EventsForWeek(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	if len(user.EventsForWeek()) == 0 {
		return nil, errors.New("Events not found ")
	}

	return user.EventsForWeek(), nil
}

func (service *GetService) EventsForMonth(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	if len(user.EventsForMonth()) == 0 {
		return nil, errors.New("Events not found ")
	}

	return user.EventsForMonth(), nil
}
