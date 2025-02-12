package grpcServer

import (
	"context"
	"errors"
	"github.com/apix76/ShortenURL/Db/DbInterface"
	"github.com/apix76/ShortenURL/Proto"
	"github.com/apix76/ShortenURL/UseCase"
)

type Server struct {
	Proto.UnimplementedShortenURLServer
	Db DbInterface.Db
}

func (s *Server) GetShortenURL(ctx context.Context, url *Proto.URL) (*Proto.ShortURL, error) {
	if url == nil {
		return nil, errors.New("url can't be nil")
	}

	shortUrl := Proto.ShortURL{ShortURL: UseCase.ShortenURL(url.Url)}

	if _, err := s.Db.Get(shortUrl.ShortURL); err != nil {
		if err == DbInterface.ErrNoExist {
			if err = s.Db.Add(shortUrl.ShortURL, url.Url); err != nil {
				return nil, err
			}
		} else {
			return &shortUrl, err
		}
	}

	return &shortUrl, nil
}

func (s Server) GetAllURL(ctx context.Context, shortUrl *Proto.ShortURL) (*Proto.URL, error) {
	if shortUrl == nil {
		return nil, errors.New("shortenUrl cannot be nil")
	}

	var (
		url Proto.URL
		err error
	)

	if url.Url, err = s.Db.Get(shortUrl.ShortURL); err != nil {
		return nil, err
	}

	return &url, err
}
