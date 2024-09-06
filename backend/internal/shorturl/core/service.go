package core

type ShortUrlService interface {
	CreateUrl(url *UrlRequest) (*Url, error)
	GetUrl(id string) (*Url, error)
}
