package models

import (
	"fmt"
	"time"

	"example.com/go_rest_api/db"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	insert into events("name", "description", "location", "datetime", "user_id")
	values ($1, $2, $3, $4, $5)
	returning id
	`
	var returningId int64
	err := db.DB.QueryRow(
		query,
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.UserId,
	).Scan(&returningId)

	e.Id = returningId

	return err

}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {

	query := "select * from events where id = $1"
	var e Event

	row := db.DB.QueryRow(query, id)
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		fmt.Println(err)
		return e, err
	}

	return e, nil
}

func (e Event) Update() error {
	query := `
	update events
	set name=$1, description=$2, location=$3, datetime=$4
	where id=$5
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

func (e Event) Delete() error {
	query := `delete from events where id=$1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)
	return err
}

func (e Event) Register(uId int64) error {
	query := "insert into registrations (user_id, event_id) values ($1, $2) returning id"
	var returningId int64
	err := db.DB.QueryRow(
		query,
		uId,
		e.Id,
	).Scan(&returningId)

	return err
}

func (e Event) CancelRegistration(uId int64) error {
	query := "delete from registrations where user_id=$1 and event_id=$2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uId, e.Id)
	return err
}
