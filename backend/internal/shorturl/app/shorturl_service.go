package app

import (
	"fmt"
	"time"
	"url-shortener/internal/shared/util"
	"url-shortener/internal/shorturl/core"

	"github.com/teris-io/shortid"
)

type ShortUrlService struct {
	repo                  core.ShortUrlRepository
	originalUrlIndex      string
	originalUrlIndexField string
}

func NewShortUrlService(repo core.ShortUrlRepository, originalUrlIndex string, originalUrlIndexField string) *ShortUrlService {
	return &ShortUrlService{repo: repo, originalUrlIndex: originalUrlIndex, originalUrlIndexField: originalUrlIndexField}
}

func generateShortID() string {
	id, _ := shortid.Generate()
	return id
}

func (s *ShortUrlService) CreateUrl(UrlRq *core.UrlRequest) (*core.Url, error) {
	urlExists := s.findOriginalUrl(UrlRq.UrlOriginal)
	if urlExists != nil {
		return urlExists, nil
	}
	id := generateShortID()

	urlSave := core.Url{
		ID:          id,
		UrlShorted:  fmt.Sprintf("%s%s", util.DOMAIN, id),
		UrlOriginal: UrlRq.UrlOriginal,
		CreatedAt:   time.Now().String(),
	}

	if err := s.repo.Save(&urlSave); err != nil {
		return nil, err
	}

	return &urlSave, nil
}
func (s *ShortUrlService) findOriginalUrl(originalUrl string) *core.Url {

	urls, err := s.repo.GetByField(s.originalUrlIndex, s.originalUrlIndexField, originalUrl)

	if err != nil || urls == nil {
		return nil
	}
	return urls[0]
}

func (s *ShortUrlService) GetUrl(id string) (*core.Url, error) {
	return s.repo.GetByID(id)
}
