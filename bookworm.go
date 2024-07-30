package main

import "os"

// loadBookworms reads the file and returns the list of
// bookworms, and their beloved books, found therein.
func loadBookworms(filePath string) ([]Bookworm, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f, nil

	// return nil, nil

}
