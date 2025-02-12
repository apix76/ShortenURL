package InMemory

import (
	"github.com/apix76/ShortenURL/Db/DbInterface"
	"sync"
)

type DbMap struct {
	db map[string]string
	mu sync.Mutex
}

func NewDb() DbMap {
	db := DbMap{db: make(map[string]string)}
	return db
}

func (db *DbMap) Add(shortURL, URL string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, ok := db.db[shortURL]
	if ok {
		err := DbInterface.ErrNoExist
		return err
	}

	db.db[shortURL] = URL

	return nil
}

func (db *DbMap) Get(shortURL string) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	URL, ok := db.db[shortURL]
	if !ok {
		err := DbInterface.ErrNoExist
		return "", err
	}

	return URL, nil
}

func (db *DbMap) Delete(shortURL string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.db, shortURL)
	return nil
}
