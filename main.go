package main

import (
	"os"

	"github.com/kyrylyuk-andriy/golang-homework/hw4"
)

func main() {
	lines := hw4.ReadLines(os.Stdin)
	hw4.PrintLines(lines, os.Stdout)
}
