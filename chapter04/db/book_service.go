package db

import "errors"

type Book struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	OwnerID string `json:"owner_id"`
	Status  string `json:"status"`
}

type BookService struct {
	books map[string]Book
	ps    PostingService
}

func NewBookService(initial []Book, ps PostingService) *BookService {
	books := make(map[string]Book)
	for _, b := range initial {
		books[b.ID] = b
	}
	return &BookService{
		books: books,
		ps:    ps,
	}
}

func (bs *BookService) Get(id string) (*Book, error) {
	b, ok := bs.books[id]
	if !ok {
		return nil, errors.New("no book found")
	}

	return &b, nil
}

// func (bs *BookService) Upsert(b Book) Book
//
// func (bs *BookService) List() []Book
//
// func (bs *BookService) ListByUser(userID string) []Book
//
// func (bs *BookService) SwapBook(bookID, userID string) (*Book, error)
