package Db

import (
	"github.com/apix76/ShortenURL/Db/DbInterface"
	"github.com/apix76/ShortenURL/Db/InMemory"
	"github.com/apix76/ShortenURL/Db/Psql"
)

func NewDb(dsn string) (DbInterface.Db, error) {
	var Db DbInterface.Db

	if dsn != "" {
		NewDb, err := Psql.NewDb(dsn)
		if err != nil {
			return nil, err
		}
		Db = &NewDb
	} else {
		NewDb := InMemory.NewDb()
		Db = &NewDb
	}

	return Db, nil
}
