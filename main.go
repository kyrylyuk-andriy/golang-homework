package main

import (
	"os"

	"github.com/kyrylyuk-andriy/golang-homework/homework"
)

func main() {
	lines := homework.ReadLines(os.Stdin)
	homework.PrintLines(lines, os.Stdout)
}
