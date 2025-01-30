package Repos

import (
	Domain2 "WB2/internal/Application/Domain"
	"database/sql"
)

type UserEventsRepo struct {
	db *sql.DB
}

func NewUserEventsRepo(db *sql.DB) *UserEventsRepo {
	return &UserEventsRepo{db: db}
}

func (repo *UserEventsRepo) GetUserById(userId int) (*Domain2.User, error) {
	var user Domain2.User
	var key string

	query := `SELECT Id, Key FROM Users WHERE Id = $1`
	err := repo.db.QueryRow(query, userId).Scan(&user.Id, &key)
	if err != nil {
		return nil, err
	}
	user.Key = key

	events, err := repo.GetEventsByUserId(userId)
	if err != nil {
		return nil, err
	}
	user.Events = events

	return &user, nil
}

func (repo *UserEventsRepo) GetEventsByUserId(userId int) ([]Domain2.Event, error) {
	var events []Domain2.Event

	query := `SELECT e.Id, e.Date, e.Description 
              FROM Events e 
              JOIN UserEvents ue ON e.Id = ue.EventId 
              WHERE ue.UserId = $1`
	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return []Domain2.Event{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Domain2.Event
		if err := rows.Scan(&event.Id, &event.Date, &event.Description); err != nil {
			return []Domain2.Event{}, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return []Domain2.Event{}, err
	}

	return events, nil
}

func (repo *UserEventsRepo) UpdateUser(user *Domain2.User) error {
	query := `UPDATE Users SET Key = $1 WHERE Id = $2`
	_, err := repo.db.Exec(query, user.Key, user.Id)
	if err != nil {
		return err
	}

	for _, event := range user.Events {
		updateEventQuery := `UPDATE Events SET Date = $1, Description = $2 WHERE Id = $3`
		_, err := repo.db.Exec(updateEventQuery, event.Date, event.Description, event.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *UserEventsRepo) AddEvent(userId int, event *Domain2.Event) error {
	query := `INSERT INTO Events (Date, Description) VALUES ($1, $2) RETURNING Id`
	err := repo.db.QueryRow(query, event.Date, event.Description).Scan(&event.Id)
	if err != nil {
		return err
	}

	userEventQuery := `INSERT INTO UserEvents (UserId, EventId) VALUES ($1, $2)`
	_, err = repo.db.Exec(userEventQuery, userId, event.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserEventsRepo) AddUser(user *Domain2.User) error {
	query := `INSERT INTO Users (Key) VALUES ($1) RETURNING Id`
	err := repo.db.QueryRow(query, user.Key).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserEventsRepo) DeleteEvent(userId int, eventId int) error {
	deleteUserEventQuery := `DELETE FROM UserEvents WHERE UserId = $1 AND EventId = $2`
	_, err := repo.db.Exec(deleteUserEventQuery, userId, eventId)
	if err != nil {
		return err
	}

	deleteEventQuery := `DELETE FROM Events WHERE Id = $1`
	_, err = repo.db.Exec(deleteEventQuery, eventId)
	if err != nil {
		return err
	}

	return nil
}
