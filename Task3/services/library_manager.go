package services

import "github.com/Johna210/backend_assessment/Track3/models"

type LibaryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks(memberID int) []models.Book
}