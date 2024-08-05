package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Johna210/backend_assessment/Track3/models"
	"github.com/Johna210/backend_assessment/Track3/services"
)

type Book = models.Book
type Member =  models.Member

type Library = services.Library

// AddBookController is a function that adds a book to the library.
// It takes a pointer to a Library struct and a pointer to a bufio.Reader as parameters.
// The function prompts the user to enter the book ID, title, and author using the reader.
// It then creates a Book struct with the entered information and a default status of "Available".
// The function calls the AddBook method of the library to add the book.
// If an error occurs during the addition, it prints an error message.
// Otherwise, it prints a success message.
func  AddBookController(library *Library, reader *bufio.Reader) {
	fmt.Println("Enter Book ID: ")	
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Println("Enter Book Title: ")
	titleStr,_ := reader.ReadString('\n')
	title := strings.TrimSpace(titleStr)

	fmt.Println("Enter Book Author: ")
	authorStr,_ := reader.ReadString('\n')
	author := strings.TrimSpace(authorStr)

	book := Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	err := library.AddBook(book)
	if err != nil {
		fmt.Println("Error adding book:", err)
	} else {
		fmt.Println("Book added successfully.")
	}

}

// RemoveBookController removes a book from the library based on the provided book ID.
// It takes a pointer to a Library object and a reader for user input as parameters.
// The function prompts the user to enter a book ID, removes the book from the library,
// and prints a success message if the book is removed successfully.
// If an error occurs during the removal process, an error message is printed.
func RemoveBookController(library *Library, reader *bufio.Reader) {
	fmt.Println("Enter Book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, _ := strconv.Atoi(bookIDStr)

	err := library.RemoveBook(bookID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

// BorrowBookController handles the borrowing of a book from the library.
// It prompts the user to enter the book ID and member ID, and then calls the BorrowBook method of the library.
// If the borrowing is successful, it prints a success message. Otherwise, it prints an error message.
func BorrowBookController(library *Library, reader *bufio.Reader) {
	fmt.Println("Enter Book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, _ := strconv.Atoi(bookIDStr)

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}


// ReturnBookController is a function that handles the process of returning a book from the library.
// It takes a pointer to a Library object and a pointer to a bufio.Reader object as parameters.
// The function prompts the user to enter the book ID and member ID, and then calls the ReturnBook method of the Library object
// to return the book. If an error occurs during the process, it prints an error message. Otherwise, it prints a success message.
func ReturnBookController(library *Library, reader *bufio.Reader){
	fmt.Println("Enter Book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, _ := strconv.Atoi(bookIDStr)

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	err := library.ReturnBook(bookID,memberID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book Returned successfully")
	}

}


// ListAvailableBooksController lists all the available books in the library.
func ListAvailableBooksController(library *Library, reader *bufio.Reader){
	availableBooks := library.ListAvailableBooks()
	fmt.Println(availableBooks)
	
}

func ListBorrowedBooksController(library *Library, reader *bufio.Reader){
	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	borrowedBooks := library.ListBorrowedBooks(memberID)
	fmt.Println(borrowedBooks)
}