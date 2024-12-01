package provider

import (
	"database/sql"
	"fmt"
	"log"
)

type DatabaseProvider struct {
	db *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *DatabaseProvider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &DatabaseProvider{db: conn}
}

func (dp *DatabaseProvider) SelectUser(name string) (string, error) {
	var user string
	row := dp.db.QueryRow("SELECT name FROM mytable WHERE name = $1", name)
	err := row.Scan(&user)
	if err != nil {
		return "", err
	}
	return user, nil
}

func (dp *DatabaseProvider) InsertUser(name string) error {
	_, err := dp.db.Exec("INSERT INTO mytable (name) VALUES ($1)", name)
	return err
}
