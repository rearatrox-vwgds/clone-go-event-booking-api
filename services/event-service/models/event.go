package models

import (
	"fmt"
	"time"

	"rearatrox/event-booking-api/services/event-service/db"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	CreatorID   int64
}

var events = []Event{}

func (e *Event) SaveEvent() error {

	query := `
	INSERT INTO events (name, description, location, dateTime, creator_id)
	VALUES (?, ?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.CreatorID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) UpdateEvent() error {
	query := `UPDATE events SET name= ?, description=?, location=?,dateTime=?,creator_id=? WHERE id=?`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.CreatorID, e.ID)

	return err
}

func (e *Event) DeleteEvent() error {
	query := `DELETE from events WHERE id=?`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID)

	return err
}

func GetEvents() ([]Event, error) {

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.CreatorID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event
	query := `SELECT * FROM events where ID = ?`

	res := db.DB.QueryRow(query, id)

	err := res.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.CreatorID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO event_registrations(event_id, user_id) VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Prepare failed")
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}

func (e Event) DeleteRegistration(userId int64) error {
	query := `DELETE FROM event_registrations where event_id = ? and user_id = ?`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}
