package core

type ShortUrlRepository interface {
	Save(url *Url) error
	GetByID(id string) (*Url, error)
}
