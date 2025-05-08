package storage

import (
	"sis-pustaka/model"
	"sis-pustaka/utils"
	"sync"
)	

type bookStorage struct {
	data map[string]model.Book
	mu    sync.RWMutex
}	

// Singleton instance
var bookStore = utils.NewOnceSingleton(func() interface{} {
	return &bookStorage{
		data: make(map[string]model.Book),
	}
})

// Getter untuk singleton instance
func GetBookStorage() *bookStorage {
	return bookStore.Get().(*bookStorage)
}

// Get all books
func (bs *bookStorage) GetAll() []model.Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	books := make([]model.Book, 0, len(bs.data))
	for _, book := range bs.data {
		books = append(books, book)
	}
	return books
}

// Get book by ID
func (bs *bookStorage) GetByID(id string) (model.Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	book, found := bs.data[id]
	return book, found
}

// Create book
func (bs *bookStorage) Create(book model.Book) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.data[book.ID] = book
}

// Update book
func (bs *bookStorage) Update(id string, book model.Book) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, exists := bs.data[id]; !exists {
		return false
	}

	bs.data[id] = book
	return true
}

// Delete book
func (bs *bookStorage) Delete(id string) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, exists := bs.data[id]; !exists {
		return false
	}

	delete(bs.data, id)
	return true
}