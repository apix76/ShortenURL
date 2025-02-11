package Proto

import (
	"ReductionAPI/Db/DbInterface"
	"ReductionAPI/UseCase"
	"context"
	"database/sql"
	"errors"
)

type Server struct {
	UnimplementedShortenURLServer
	Db DbInterface.Db
}

func (s *Server) GetShortenURL(ctx context.Context, url *URL) (*ShortURL, error) {
	if url == nil {
		return nil, errors.New("url can't be nil")
	}

	shortUrl := ShortURL{ShortURL: UseCase.ShortenURL(url.Url)}

	if _, err := s.Db.Get(shortUrl.ShortURL); err != nil {
		if err == DbInterface.ErrNoExist || err == sql.ErrNoRows {
			if err = s.Db.Add(shortUrl.ShortURL, url.Url); err != nil {
				return nil, err
			}
		} else {
			return &shortUrl, err
		}
	}

	return &shortUrl, nil
}

func (s Server) GetAllURL(ctx context.Context, shortUrl *ShortURL) (*URL, error) {
	if shortUrl == nil {
		return nil, errors.New("shortenUrl cannot be nil")
	}

	var (
		url URL
		err error
	)

	if url.Url, err = s.Db.Get(shortUrl.ShortURL); err != nil {
		return nil, err
	}

	return &url, err
}
