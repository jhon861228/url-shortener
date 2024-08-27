package app

import "url-shortener/internal/shorturl/core"

type ShortUrlService struct {
	repo core.ShortUrlRepository
}

func NewShortUrlService(repo core.ShortUrlRepository) *ShortUrlService {
	return &ShortUrlService{repo: repo}
}

func (s *ShortUrlService) CreateUrl(Url *core.Url) error {
	return s.repo.Save(Url)
}

func (s *ShortUrlService) GetUrl(id string) (*core.Url, error) {
	return s.repo.GetByID(id)
}
