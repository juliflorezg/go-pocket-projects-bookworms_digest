package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load the bookworms file: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", bookworms)
	fmt.Printf("%v\n", bookworms)

	fmt.Printf("book count: %+v\n", booksCount(bookworms))
	emptyBookworms := []Bookworm{}
	var emptyBookworms2 []Bookworm
	fmt.Printf("no bookworms: %+v\n", booksCount(emptyBookworms))

	// ! go won't even let the comparison happen:
	// ! invalid operation: emptyBookworms == emptyBookworms (slice can only be compared to nil)
	// fmt.Printf("comparing empty bookworms slice to itself: %v\n", emptyBookworms == emptyBookworms)
	fmt.Printf("comparing emptyBookworms slice to nil: %v\n", emptyBookworms == nil)
	fmt.Printf("comparing emptyBookworms2 slice to nil: %v\n", emptyBookworms2 == nil)
	fmt.Println(len(emptyBookworms))
	fmt.Println(len(emptyBookworms2))
}
