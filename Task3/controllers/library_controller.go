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