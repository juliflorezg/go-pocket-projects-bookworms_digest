package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	villette      = Book{Author: "Charlotte Brontë", Title: "Villette"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantError     bool
	}

	tests := map[string]testCase{
		"File exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantError: false,
		},
		"File does not exist": {
			bookwormsFile: "testdata/no_file.json",
			want:          nil,
			wantError:     true,
		},
		"File is and invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantError:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantError {
				if err == nil {
					t.Fatal("Expected an error but got nothing")
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error but got %v", err)
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("Different results: got\n %+v\n expected\n %+v", got, tc.want)
			}
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(t *testing.T, books, targetBooks []Book) bool {
	t.Helper()

	if len(books) != len(targetBooks) {
		return false
	}

	for i := range targetBooks {
		if targetBooks[i] != books[i] {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {

	type testCase struct {
		bookworms []Bookworm
		want      map[Book]uint
		// got       map[Book]uint
	}

	tests := map[string]testCase{
		"2 exact book counts": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		// TODO: fill with more test cases
		"no bookworms": {
			bookworms: []Bookworm{},
			want:      map[Book]uint{},
		},
		"bookworm with no books": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{
				oryxAndCrake:  1,
				handmaidsTale: 1,
				janeEyre:      1,
			},
		},
		"bookworm with two of the same book": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre, handmaidsTale}},
			},
			want: map[Book]uint{
				handmaidsTale: 3,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			got := booksCount(tc.bookworms)

			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("Book count is not equal, got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]

		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestFindCommonBooks(t *testing.T) {
	type testCase struct {
		input []Bookworm
		want  []Book
	}

	tests := map[string]testCase{
		"Bookworms have read different books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, villette, janeEyre}},
			},
			want: []Book{},
		},
		"Bookworms have read the same exact books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, oryxAndCrake, villette, janeEyre}},
				{Name: "Peggy", Books: []Book{handmaidsTale, theBellJar, oryxAndCrake, villette, janeEyre}},
				{Name: "Carrot", Books: []Book{handmaidsTale, theBellJar, oryxAndCrake, villette, janeEyre}},
			},
			want: []Book{
				handmaidsTale, theBellJar, oryxAndCrake, villette, janeEyre,
			},
		},
		"More than two bookworms have a book in common": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, villette, janeEyre, theBellJar}},
				{Name: "Vivi", Books: []Book{villette, theBellJar, oryxAndCrake}},
			},
			want: []Book{
				theBellJar, oryxAndCrake, villette,
			},
		},
		"One bookworm has no books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: []Book{},
		},
		"Nobody has any books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
				{Name: "Vivi", Books: []Book{}},
				{Name: "Carrot", Books: []Book{}},
			},
			want: []Book{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)

			if !equalCommonBooks(t, got, tc.want) {
				t.Fatalf("Got different lists for common books\n got: %v \n want: %v", got, tc.want)
			}
		})
	}
}

func equalCommonBooks(t *testing.T, got, want []Book) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for _, b := range want {
		// starts by assuming the current book is not in <got []Book>, when it's found in got, switch the flag to true and continue with next book to check in <want []Book>
		isBookInGotResult := false
		for _, b2 := range got {
			if b == b2 {
				isBookInGotResult = true
				break
			}
		}

		if !isBookInGotResult {
			return false
		}
	}

	return true
}
