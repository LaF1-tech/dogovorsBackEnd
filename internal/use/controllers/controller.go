package controllers

import (
	"dogovorsBackEnd/internal/use/controllers/validators"
	"dogovorsBackEnd/internal/use/repositories"
)

type Controller interface {
}

type controller struct {
	validator validators.Validator

	repository repositories.Repository
}

func New(validator validators.Validator, repository repositories.Repository) Controller {
	return &controller{validator: validator, repository: repository}
}
