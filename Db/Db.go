package Db

import (
	"ReductionAPI/Db/DbInterface"
	"ReductionAPI/Db/InMemory"
	"ReductionAPI/Db/Psql"
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
