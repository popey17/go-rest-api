package models

import (
	"time"

	"github.com/popey17/go-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(
	name, description, location, dateTime, userId)
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAll() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	results, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for results.Next() {
		var event Event
		err := results.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?`

	result := db.DB.QueryRow(query, id)

	var event Event

	err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?
	WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

	return err

}

func (e Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}
