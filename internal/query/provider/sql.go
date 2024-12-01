// internal/query/provider/sql.go
package provider

import (
	"database/sql"
)

type Provider struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{db: db}
}

func (p *Provider) SelectUser(name string) (string, error) {
	var user string
	row := p.db.QueryRow("SELECT name FROM mytable WHERE name = $1", name)
	err := row.Scan(&user)
	if err != nil {
		return "", err
	}
	return user, nil
}

func (p *Provider) InsertUser(name string) error {
	_, err := p.db.Exec("INSERT INTO mytable (name) VALUES ($1)", name)
	return err
}
