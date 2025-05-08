package book_handler

import (
	"sis-pustaka/model"
	"strings"
)

func validateBook(book model.Book) (bool, string) {
	if strings.TrimSpace(book.ID) == "" {
		return false, "ID tidak boleh kosong"
	}
	if strings.TrimSpace(book.Title) == "" {
		return false, "Title tidak boleh kosong"
	}
	if strings.TrimSpace(book.Author) == "" {
		return false, "Author tidak boleh kosong"
	}
	if strings.TrimSpace(book.Publisher_Year) == "" {
		return false, "Publisher Year tidak boleh kosong"
	}
	return true, ""
}
