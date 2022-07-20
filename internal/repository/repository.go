package repository

import (
	"database/sql"
	"library/internal/repository/author"
)

type Repository struct {
	AuthorRepository author.AuthorRepository
}

func NewRepository(db *sql.DB) *Repository {
	a := author.NewAuthorRepository(db)
	return &Repository{
		AuthorRepository: a,
	}
}
