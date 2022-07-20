package handlers

import (
	"context"
	"library/internal/delivery/grpc/librarypb"
	"library/internal/domain/model"
	"library/internal/service/author"
)

type AuthorHandlers interface {
	GetAuthorByBook(context.Context, *librarypb.BookEntity) (*librarypb.GetAuthorByBookResponse, error)
}

type authorHandlers struct {
	AuthorService author.AuthorService
}

func (ah *authorHandlers) GetAuthorByBook(ctx context.Context, req *librarypb.BookEntity) (*librarypb.GetAuthorByBookResponse, error) {
	var book model.Books
	book.Id = req.GetId()
	book.Name = req.GetName()
	book.AuthorID = req.GetAuthorID()
	author, err := ah.AuthorService.GetAuthorByBook(&book)
}
