package Db

import (
	"github.com/apix76/ShortenURL/Db/DbInterface"
	"github.com/apix76/ShortenURL/Db/InMemory"
	"github.com/apix76/ShortenURL/Db/Psql"
)

func NewDb(dns string) (DbInterface.Db, error) {
	var Db DbInterface.Db

	if dns != "" {
		NewDb, err := Psql.NewDb(dns)
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
