package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	psqlInfo := os.Getenv("DATABASE_URL")

	var err error

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Could not conect to postgres")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	create table if not exists users(
		id serial primary key,
		email text not null unique,
		password text not null
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	create table if not exists events(
		id serial primary key,
		name text not null,
		description text not null,
		location text not null,
		dateTime timestamp,
		user_id serial references users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	create table if not exists registrations(
		id serial primary key,
		user_id serial references users(id),
		event_id serial references events(id)
	)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table")
	}
}
