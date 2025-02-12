package service

import (
	"URLShort/internal/model"
	"URLShort/internal/pkg/codeGenerator"
	"URLShort/internal/storage/sqlite"
	"fmt"
)

type Service struct {
	storage *sqlite.Storage
}

func New(storage *sqlite.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateUrl(url string) (*model.ShortUrl, error) {

	code := codeGenerator.GenerateCode()
	err := s.storage.SaveUrl(url, code)

	if err != nil {
		return nil, err
	}

	m, err := s.storage.GetUrl(code)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Service) GetUrl(code string) (*model.ShortUrl, error) {
	m, err := s.storage.GetUrl(code)

	if err != nil {
		return nil, err
	}

	err = s.storage.VisitUrl(&m.AccessCount, code)

	if err != nil {
		return nil, fmt.Errorf("Error visiting URL: %v", err)
	}
	return m, nil
}

func (s *Service) GetUrlStats(code string) (*model.ShortUrl, error) {

	m := new(model.ShortUrl)

	m, err := s.storage.GetUrlStats(code)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Service) UpdateUrl(code string, url string) (*model.ShortUrl, error) {
	err := s.storage.UpdateUrl(code, url)

	if err != nil {
		return nil, err
	}

	m, err := s.storage.GetUrl(code)

	if err != nil {
		return nil, err
	}

	return m, nil

}

func (s *Service) DeleteUrl(code string) error {
	err := s.storage.DeleteUrl(code)

	return err
}
