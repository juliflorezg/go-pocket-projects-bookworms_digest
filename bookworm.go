package main

import (
	"encoding/json"
	"os"
	"sort"
)

// A Bookworm contains the list of books on bookworm's shelf
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes book information on a bookworm's shelf
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of
// bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Initialize the type in which the file will be decoded.
	var bookworms []Bookworm

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// booksCount registers all the books and their occurrences from the bookworms shelves.
func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworms := range bookworms {
		for _, book := range bookworms.Books {
			// takes two copies of the same book as valid
			count[book]++
		}
	}

	return count
}

// findCommonBooks returns books that are on more than one bookworm's shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)

	commonBooks := make([]Book, 0)

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// sortBooks sorts the books by Author and then Title.
func sortBooks(books []Book) []Book {
	// slices.SortFunc[]()
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}

		return books[i].Title < books[j].Title
	})

	return books
}
