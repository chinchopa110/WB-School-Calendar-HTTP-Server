package UserService

import (
	"WB2/internal/Application/Abstractions/Repos"
	"WB2/internal/Application/Domain"
	"errors"
)

type GetService struct {
	repo Repos.IUserEventsRepo
}

func CreateGetService(repo Repos.IUserEventsRepo) *GetService {
	return &GetService{repo: repo}
}

func (service GetService) Authentication(userId int, key string) error {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return err
	}
	if !user.IsKey(key) {
		return errors.New("INCORRECT KEY")
	}
	return nil
}

func (service GetService) EventsForDay(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	return user.EventsForDay(), nil
}

func (service GetService) EventsForWeek(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	return user.EventsForWeek(), nil
}

func (service GetService) EventsForMonth(userId int, key string) ([]Domain.Event, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if !user.IsKey(key) {
		return nil, errors.New("INCORRECT KEY")
	}

	return user.EventsForMonth(), nil
}
