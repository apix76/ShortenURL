package Psql

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbAccess struct {
	db *sql.DB
}

func NewDb(dsn string) (DbAccess, error) {
	var Db DbAccess
	var err error

	if Db.db, err = sql.Open("pgx", dsn); err != nil {
		return Db, err
	}
	return Db, err
}

func (db *DbAccess) Add(shortURL, URL string) error {
	if _, err := db.db.Exec("insert into links (shortURL, URL) values ($1, $2)", shortURL, URL); err != nil {
		return err
	}
	return nil
}

func (db *DbAccess) Get(shortURL string) (string, error) {
	row := db.db.QueryRow("select URL from links where shortURL = $1", shortURL)

	var URL string

	if err := row.Scan(&URL); err != nil {
		return "", err
	}
	return URL, nil
}

func (db *DbAccess) Delete(shortURL string) error {
	_, err := db.db.Exec("delete from links where shortURL = $1", shortURL)
	return err
}
