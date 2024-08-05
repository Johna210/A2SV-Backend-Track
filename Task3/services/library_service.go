package services

import (
	"errors"

	"github.com/Johna210/backend_assessment/Track3/models"
)


type Book = models.Book
type Member = models.Member


type Library struct {
	Books map[int]Book
	Members map[int]Member
}

// AddBook adds a new book to the library.
// It checks if the book already exists in the library by checking its ID.
// If the book already exists, it returns an error.
// Otherwise, it adds the book to the library and returns nil.
func (l *Library) AddBook(book Book) error {
	if _, exists := l.Books[book.ID]; exists {
		return errors.New("book already exists")
	}
	l.Books[book.ID] = book
	return nil
}

// RemoveBook removes a book from the library based on the given bookID.
// If the book is not found, it returns an error.
func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

// BorrowBook borrows a book from the library for a given member.
// It updates the book's status to "Borrowed" and adds the book to the member's borrowed books list.
// If the book or member is not found, or if the book is already borrowed, an error is returned.
func (l *Library ) BorrowBook(bookID int,memberID int) error {
	book, bookExists := l.Books[bookID]
	member, memeberExists := l.Members[memberID]

	if !bookExists {
		return errors.New("book not found")
	}

	if !memeberExists {
		return errors.New("memeber not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

// ReturnBook returns a borrowed book to the library.
// It takes the bookID and memberID as parameters and returns an error if any.
// If the book is not found, it returns an error with the message "book not found".
// If the member is not found, it returns an error with the message "member not found".
// If the book is not borrowed, it returns an error with the message "book was not borrowed".
// If the book is not borrowed by the specified member, it returns an error with the message "book not borrowed by this member".
// Otherwise, it updates the book status to "Available" and removes the book from the member's borrowed books list.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	if book.Status != "Borrowed" {
		return errors.New("book was not borrowed")
	}

	bookFound := false

	for _,borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == book.ID {
			bookFound = true
			break
		}
	}

	if !bookFound {
		return errors.New("book not borrowed by this member")

	}

	// Update the book status to "Available"
	book.Status = "Available"
	l.Books[bookID] = book

	// Remove the book from the member's borrowed books list
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	return nil
}

// ListAvailableBooks returns a list of available books in the library.
func (l *Library) ListAvailableBooks() []Book {
	allBooks := make([]Book,0)

	for _,book := range l.Books {
		if book.Status == "Available"{
			allBooks = append(allBooks,book)
		}
	}

	return allBooks
}

// ListBorrowedBooks returns a list of books borrowed by a member with the specified memberID.
// If the member does not exist, an empty list is returned.
func (l *Library) ListBorrowedBooks(memberID int) []Book {
	member, exists := l.Members[memberID]

	if !exists {
		return []Book{}
	}


	return member.BorrowedBooks
}