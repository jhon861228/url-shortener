package core

type ShortUrlService interface {
	CreateUrl(url *Url) error
	GetUrl(id string) (*Url, error)
}
