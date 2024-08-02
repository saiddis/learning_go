package main

type Book struct {
	title  string
	author string
	year   int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(title, author string, year int) {
	book := Book{
		title:  title,
		author: author,
		year:   year,
	}

	l.Books = append(l.Books, book)
}

func (l Library) BooksByAuthorAfterYear(year int) []Book {
	var booksAfterYear []Book

	for _, v := range l.Books {
		if v.year > year {
			booksAfterYear = append(booksAfterYear, v)
		}
	}

	return booksAfterYear
}
