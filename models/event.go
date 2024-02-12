package models

import (
	"log"
	"myapp/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
    INSERT INTO events (name, description, location, dateTime, user_id) 
    OUTPUT INSERTED.id
    VALUES (@p1, @p2, @p3, @p4, @p5);
    `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&id)
	if err != nil {
		log.Printf("Error executing query and getting last insert ID: %v", err)
		return err
	}

	e.Id = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events
	WHERE id = @p1
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = @p1, description = @p2, location = @p3, dateTime = @p4
	WHERE id = @p5
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}

func (event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = @p1
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id)
	return err
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id)
	VALUES(@p1, @p2)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE event_id = @p1 AND user_id = @p2
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)

	return err
}
