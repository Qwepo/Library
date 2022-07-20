package service

import (
	"library/internal/repository"
	"library/internal/service/author"
)

type Service struct {
	Author author.AuthorService
}

func NewService(r *repository.Repository) *Service {
	a := author.NewAuthrService(r.AuthorRepository)
	return &Service{
		Author: a,
	}
}
