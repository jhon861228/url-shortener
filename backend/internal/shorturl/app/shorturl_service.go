package app

import (
	"fmt"
	"time"
	"url-shortener/internal/shared/util"
	"url-shortener/internal/shorturl/core"

	"github.com/google/uuid"
)

type ShortUrlService struct {
	repo core.ShortUrlRepository
}

func NewShortUrlService(repo core.ShortUrlRepository) *ShortUrlService {
	return &ShortUrlService{repo: repo}
}

func (s *ShortUrlService) CreateUrl(UrlRq *core.UrlRequest) (*core.Url, error) {
	id := uuid.New()
	urlSave := core.Url{
		ID:          id.String(),
		UrlShorted:  fmt.Sprintf("%s%s", util.DOMAIN, id),
		UrlOriginal: UrlRq.UrlOriginal,
		CreatedAt:   time.Now().String(),
	}

	if err := s.repo.Save(&urlSave); err != nil {
		return nil, err
	}

	return &urlSave, nil
}

func (s *ShortUrlService) GetUrl(id string) (*core.Url, error) {
	return s.repo.GetByID(id)
}
