package data

import "strings"

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type DataList struct {
	LastID int
	Books  []Book
}

var Data = DataList{LastID: 0, Books: []Book{}}

func (d *DataList) AddBook(name, author string) Book {
	d.LastID++
	book := Book{Id: d.LastID, Name: name, Author: author}
	d.Books = append(d.Books, book)
	return book
}

func (d *DataList) GetBooks() []Book {
	return d.Books
}

func (d *DataList) GetBook(id int) Book {
	for _, book := range d.Books {
		if book.Id == id {
			return book
		}
	}
	return Book{}
}

func (d *DataList) DeleteBook(id int) {
	for i, book := range d.Books {
		if book.Id == id {
			d.Books = append(d.Books[:i], d.Books[i+1:]...)
			return
		}
	}
}

func (d *DataList) UpdateBook(id int, name, author string) {
	for i, book := range d.Books {
		if book.Id == id {
			d.Books[i].Name = name
			d.Books[i].Author = author
			return
		}
	}
}

func (d *DataList) SearchBooks(term string) []Book {
	//search if trem is part of name or author case insensitive
	var results []Book

	for _, book := range d.Books {
		if strings.Contains(strings.ToLower(book.Name), strings.ToLower(term)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(term)) {
			results = append(results, book)
		}
	}
	return results
}

func init() {
	Data.AddBook("The Catcher in the Rye", "J. D. Salinger")
	Data.AddBook("The Great Gatsby", "F. Scott Fitzgerald")
	Data.AddBook("Pride and Prejudice", "Jane Austen")
	Data.AddBook("The Hobbit", "J. R. R. Tolkien")
	Data.AddBook("The Lord of the Rings", "J. R. R. Tolkien")
	Data.AddBook("The Catcher in the Rye", "J. D. Salinger")
	Data.AddBook("The Alchemist", "Paulo Coelho")
	Data.AddBook("The Hobbit", "J. R. R. Tolkien")
}
