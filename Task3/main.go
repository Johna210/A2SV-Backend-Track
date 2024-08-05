package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	controllers "github.com/Johna210/backend_assessment/Track3/controllers"
	"github.com/Johna210/backend_assessment/Track3/models"
	"github.com/Johna210/backend_assessment/Track3/services"
)

type Book = models.Book
type Member =  models.Member

var library = services.Library{
	Books: make(map[int]Book),
	Members: map[int]Member{
		1: {
			ID: 1,
			Name: "John",
			BorrowedBooks: make([]Book,0),
		},
		2: {
			ID: 2,
			Name: "James",
			BorrowedBooks: make([]Book,0),
		},
		3: {
			ID: 3,
			Name: "Jack",
			BorrowedBooks: make([]Book,0),
		},
	},
}


func main()  {
	
		
	var reader = bufio.NewReader(os.Stdin)

	for {
		menu()
		fmt.Println("Enter Choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)



		switch choice {
		case 1:
			controllers.AddBookController(&library, reader)
		case 2:
			controllers.RemoveBookController(&library, reader)
		case 3:
			controllers.BorrowBookController(&library, reader)
		case 4:
			controllers.ReturnBookController(&library, reader)
		case 5:
			controllers.ListAvailableBooksController(&library, reader)
		case 6:
			controllers.ListBorrowedBooksController(&library, reader)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		
		}

		fmt.Println()
	}
	

	
}

func menu() {
	fmt.Println("Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}