package core

type ShortUrlRepository interface {
	Save(url *Url) error
	GetByID(id string) (*Url, error)
	GetByField(indexName string, fieldName string, fieldValue string) ([]*Url, error)
}
