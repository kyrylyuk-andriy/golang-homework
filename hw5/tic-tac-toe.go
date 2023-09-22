package hw5

import (
	"fmt"
)

var gameBoard = make(map[int][string])

func CreateBoardGame() {
	for i := 1; i <= 9; i++ {
		gameBoard[i] = " "
	}
}

func PrintBoardGame() {
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---|---|---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---|---|---")
	fmt.Println(" 7 | 8 | 9 ")
}