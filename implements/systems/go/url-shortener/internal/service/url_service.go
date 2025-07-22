package service

import (
	"github.com/JawherKl/url-shortener/internal/model"
	"github.com/JawherKl/url-shortener/internal/utils"
	"os"
)

type Redis interface {
	Set(string, string) error
	Get(string) (string, error)
}

type Postgres interface {
	SaveURL(model.URL) error
	GetOriginalURL(string) (string, error)
}

type URLService struct {
	redis Redis
	pg    Postgres
}

func NewURLService(r Redis, p Postgres) *URLService {
	return &URLService{redis: r, pg: p}
}

func (s *URLService) Shorten(original string) (string, error) {
	hash := utils.GenerateHash(original)
	shortURL := os.Getenv("BASE_URL") + hash

	s.redis.Set(hash, original)
	s.pg.SaveURL(model.URL{Original: original, Short: hash})

	return shortURL, nil
}

func (s *URLService) Resolve(hash string) (string, error) {
	if original, err := s.redis.Get(hash); err == nil {
		return original, nil
	}
	return s.pg.GetOriginalURL(hash)
}