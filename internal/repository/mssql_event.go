package repository

import (
	"database/sql"
	"errors"
	"log"
	"myapp/internal/models"
)

type EventRepository interface {
	Save(event *models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetEventById(id int64) (*models.Event, error)
	Update(event models.Event) error
	Delete(id int64) error
	Register(eventId, userId int64) error
	CancelRegistration(eventId, userId int64) error
}

type eventRepositoryImpl struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepositoryImpl{db}
}

func (repo *eventRepositoryImpl) Save(event *models.Event) error {
	query := `
    INSERT INTO events (name, description, location, dateTime, user_id) 
	OUTPUT INSERTED.id
    VALUES (@p1, @p2, @p3, @p4, @p5)
    `
	err := repo.db.QueryRow(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&event.Id)
	if err != nil {
		log.Printf("Error saving event: %v", err)
		return err
	}
	return nil
}

func (repo *eventRepositoryImpl) GetAllEvents() ([]models.Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (repo *eventRepositoryImpl) GetEventById(id int64) (*models.Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events WHERE id = @p1`
	event := models.Event{}
	err := repo.db.QueryRow(query, id).Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("event not found")
		}
		log.Printf("Error getting event by ID: %v", err)
		return nil, err
	}
	return &event, nil
}

func (repo *eventRepositoryImpl) Update(event models.Event) error {
	query := `
    UPDATE events
    SET name = @p1, description = @p2, location = @p3, dateTime = @p4
    WHERE id = @p5
    `
	_, err := repo.db.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.Id)
	if err != nil {
		log.Printf("Error updating event: %v", err)
		return err
	}
	return nil
}

func (repo *eventRepositoryImpl) Delete(id int64) error {
	query := `DELETE FROM events WHERE id = @p1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting event: %v", err)
		return err
	}
	return nil
}

func (repo *eventRepositoryImpl) Register(eventId, userId int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (@p1, @p2)`
	_, err := repo.db.Exec(query, eventId, userId)
	if err != nil {
		log.Printf("Error registering user to event: %v", err)
		return err
	}
	return nil
}

func (repo *eventRepositoryImpl) CancelRegistration(eventId, userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = @p1 AND user_id = @p2`
	_, err := repo.db.Exec(query, eventId, userId)
	if err != nil {
		log.Printf("Error cancelling registration: %v", err)
		return err
	}
	return nil
}
