package core

import (
	"github.com/go-playground/validator/v10"
)

type Url struct {
	ID          string `json:"id" validate:"required"`
	UrlShorted  string `json:"urlShorted" validate:"required"`
	UrlOriginal string `json:"urlOriginal" validate:"required"`
	CreatedAt   string `json:"createdAt" validate:"required"`
}

type UrlRequest struct {
	UrlOriginal string `json:"urlOriginal" validate:"required"`
}

type UrlResponse struct {
	UrlShorted string `json:"urlShorted" validate:"required"`
	CreatedAt  string `json:"createdAt" validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *Url) Validate() error {
	return validate.Struct(u)
}

func (u *UrlRequest) ValidateUrlRequest() error {
	return validate.Struct(u)
}
