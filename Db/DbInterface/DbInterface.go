package DbInterface

import (
	"errors"
)

var (
	ErrNoExist = errors.New("no exist")
)

type Db interface {
	Add(shortURL, URL string) error
	Get(shortURL string) (string, error)
	Delete(shortURL string) error
}
