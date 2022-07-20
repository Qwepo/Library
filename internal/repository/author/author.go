package author

import (
	"database/sql"
	"fmt"
	"library/internal/domain/model"
)

const (
	authorTable = "authro"
)

type AuthorRepository interface {
	Create(*model.Author) error
	GetAuthorByBook(*model.Books) (*model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}

func (ar *authorRepository) GetAuthorByBook(b *model.Books) (*model.Author, error) {
	rowb := ar.db.QueryRow(fmt.Sprintf("SELECT author_id FORM books WHERE name = %s ", b.Name))
	err := rowb.Scan(&b.AuthorID)
	if err != nil {
		return nil, err
	}
	var author model.Author
	rowa := ar.db.QueryRow(fmt.Sprintf("SELECT id, name FORM author WHERE id = %d ", b.AuthorID))
	err = rowa.Scan(&author.Id, &author.Name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (ar *authorRepository) Create(a *model.Author) error {
	insert, err := ar.db.Query(fmt.Sprintf("INSERT INTO %s (name) VALUES(%s)", authorTable, a.Name))
	if err != nil {
		return err
	}
	insert.
	defer insert.Close()
	return nil
}
