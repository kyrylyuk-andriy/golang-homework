package main

import (
	"fmt"
	"os"

	"github.com/kyrylyuk-andriy/golang-homework/hw4"
)

func main() {
	lines := hw4.ReadLines(os.Stdin)
	// optionally to print all strings to compare search result
	hw4.PrintLines(lines, os.Stdout)
	var searchQuery string
	fmt.Print("What is your search query?")
	fmt.Scan(&searchQuery)
	hw4.SearchLines(lines, searchQuery)
}
