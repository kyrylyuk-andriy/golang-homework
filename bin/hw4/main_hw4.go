package main

import (
	"fmt"
	"os"

	"github.com/kyrylyuk-andriy/golang-homework/hw4"
)

func main() {
	//
	// HW4 task1
	//
	lines := hw4.ReadLines(os.Stdin)
	hw4.PrintLines(lines, os.Stdout)
	var searchQuery string
	fmt.Print("What is your search query?")
	fmt.Scan(&searchQuery)
	hw4.SearchLines(lines, searchQuery)

	//
	//HW4 task2
	hw4.Average_score()
}
