package main

import "fmt"

type Book struct {
	ID        int
	Title     string
	Author    string
	Available bool
}

type Library struct {
	books []Book
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) ListBook() []Book {
	return l.books
}

func main() {

	manager := Library{}

	manager.AddBook(
		Book{
			ID:        1,
			Title:     "Clean Code",
			Author:    "Robert C. Martin",
			Available: true,
		},
	)

	manager.AddBook(
		Book{
			ID:        2,
			Title:     "The Pragmatic Programmer",
			Author:    "Andrew Hunt e David Thomas",
			Available: true,
		},
	)

	manager.AddBook(
		Book{
			ID:        3,
			Title:     "Design Patterns",
			Author:    "Erich Gamma, Richard Helm, Ralph Johnson e John Vlissides",
			Available: false,
		},
	)

	manager.AddBook(
		Book{
			ID:        4,
			Title:     "Entendendo Algoritmos",
			Author:    "Aditya Bhargava",
			Available: false,
		},
	)

	manager.AddBook(
		Book{
			ID:        5,
			Title:     "Go Programming Language",
			Author:    "Alan Donovan e Brian Kernighan",
			Available: true,
		},
	)

	for _, book := range manager.ListBook() {
		fmt.Printf("ID: %d\n Titile: %s\n Author: %s\n Avalaible: %t\n",
			book.ID,
			book.Title,
			book.Author,
			book.Available,
		)
	}

}
