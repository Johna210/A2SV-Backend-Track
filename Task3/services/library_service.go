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

func (l *Library) AddBook(book Book) error {
	if _, exists := l.Books[book.ID]; exists {
		return errors.New("book already exists")
	}
	l.Books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

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

func (l *Library) ListAvailableBooks() []Book {
	allBooks := make([]Book,0)

	for _,book := range l.Books {
		if book.Status == "Available"{
			allBooks = append(allBooks,book)
		}
	}

	return allBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []Book {
	member, exists := l.Members[memberID]

	if !exists {
		return []Book{}
	}


	return member.BorrowedBooks
}