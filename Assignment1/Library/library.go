package main

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			book.IsBorrowed = true
			l.Books[id] = book
			fmt.Println("You have borrowed:", book.Title)
		} else {
			fmt.Println("This book is already borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			book.IsBorrowed = false
			l.Books[id] = book
			fmt.Println("You have returned:", book.Title)
		} else {
			fmt.Println("This book was not borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for id, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", id, book.Title, book.Author)
		}
	}
}

func main() {
	library := Library{
		Books: make(map[string]Book),
	}

	for {
		var action string
		fmt.Println("\nChoose an action: Add, Borrow, Return, List, Exit")
		fmt.Scanln(&action)

		switch action {
		case "Add":
			var id, title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter book title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter book author: ")
			fmt.Scanln(&author)

			book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			library.AddBook(book)
			fmt.Println("Book added.")
		case "Borrow":
			var id string
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scanln(&id)
			library.BorrowBook(id)
		case "Return":
			var id string
			fmt.Print("Enter book ID to return: ")
			fmt.Scanln(&id)
			library.ReturnBook(id)
		case "List":
			library.ListBooks()
		case "Exit":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid action.")
		}
	}
}
