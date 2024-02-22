package main

import "os"

func main() {
	if len(os.Args) > 1 {
		print(os.Args[1])
	} else {
		print("No Args")
	}
	// - indexing
	//   - stemmer
	//   - model
	//
	// - (local) server
	os.Exit(0) // status OK
}
