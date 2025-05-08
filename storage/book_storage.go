package storage

import (
	"sis-pustaka/model"
	"sis-pustaka/utils"
	"sync"
)

	type bookStorage struct {
		data map[string]model.Book
		mu  sync.RWMutex
	}

	//singleton instance
	var bookStore = utils.NewOneSingleton(func() interface{} {
		return &bookStorage{
			data: make(map[string]model.Book),
		}
	})

	//Getter untuk singleton instance
	func GetBookStorage() *bookStorage {
		return bookStore.Get().(*bookStorage)
	}

	//Get all books
	func (s *bookStorage) GetAll() []model.Book {
		s.mu.RLock()
		defer s.mu.RUnlock()

		books := make([]model.Book, 0, len(s.data))
		for _, b := range s.data {
			books = append(books, b)
		}
		return books
	}

	//Get book by ID
	func (s *bookStorage) GetByID(id string) (model.Book, bool) {
		s.mu.RLock()
		defer s.mu.RUnlock()

		b, found := s.data[id]
		return b, found
	}

	//Create new book
	func (s *bookStorage) Create(b model.Book) {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.data[b.ID] = b
	}

	//Update book by ID
	func (s *bookStorage) Update(id string, updated model.Book) bool {
		s.mu.Lock()
		defer s.mu.Unlock()

		if _, found := s.data[id]; !found {
			return false
		}
		s.data[id] = updated
		return true
	}

	//Delete book by ID
	func (s *bookStorage) Delete(id string) bool {
		s.mu.Lock()
		defer s.mu.Unlock()

		if _, found := s.data[id]; !found {
			return false
		}
		delete(s.data, id)
		return true
	}