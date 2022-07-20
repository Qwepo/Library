package author

import (
	"library/internal/domain/model"
	"library/internal/repository/author"
)

type AuthorService interface {
	GetAuthorByBook(*model.Books) (*model.Author, error)
}

type authorService struct {
	AuthorRepository author.AuthorRepository
}

func NewAuthrService(r author.AuthorRepository) AuthorService {
	return &authorService{AuthorRepository: r}
}

func (s *authorService) GetAuthorByBook(*model.Books) (*model.Author, error) {
	
}
