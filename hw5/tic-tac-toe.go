package hw5

import (
	"fmt"
	"strconv"
)

var GameBoard = make(map[int]string)
var Player string

func CreateBoardGame() {
	for i := 1; i <= 9; i++ {
		GameBoard[i] = " "
	}
}

func PrintInstructionBoard() {
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---|---|---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---|---|---")
	fmt.Println(" 7 | 8 | 9 ")
}

func PrintBoardGame() {
	fmt.Println("---------")
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s| %s |%s \n", GameBoard[(i*3+1)], GameBoard[(i*3+2)], GameBoard[(i*3+3)])
		if i < 3 {
			fmt.Println("---------")
		}
	}
	fmt.Println()
}

func checkWhoWin() bool {
	winningCombo := [8][3]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, // Each Row
		{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, // Each Columns
		{1, 5, 9}, {3, 5, 7}, // Both Diagonals
	}

	for _, combo := range winningCombo {
		if GameBoard[combo[0]] == Player && GameBoard[combo[1]] == Player && GameBoard[combo[2]] == Player {
			return true
		}
	}

	return false
}

func isGameBoardFull() bool {
	for i := 1; i <= 9; i++ {
		if GameBoard[i] == " " {
			return false
		}
	}
	return true
}

func Tic_Tac_Toe() {
	CreateBoardGame()
	var Player = "X"
	var gameActive bool = true

	fmt.Println("Welcome to Tic-Tac-Toe!")
	PrintInstructionBoard()
	fmt.Println("Each cell in a table represents to a number")
	fmt.Println("Use numbers 1-9 to make your move.")
	PrintBoardGame()

	for gameActive {
		fmt.Printf("Player %s's turn.\n", Player)
		fmt.Print("Enter your move: ")

		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 9.")
			continue
		}

		position, err := strconv.Atoi(input)

		if err != nil || position < 1 || position > 9 || GameBoard[position] != " " {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		GameBoard[position] = Player
		PrintBoardGame()

		if checkWhoWin() {
			fmt.Printf("Player %s wins!\n", Player)
			gameActive = false
		} else if isGameBoardFull() {
			fmt.Println("It's a tie!")
			gameActive = false
		} else {
			// Switch player
			if Player == "X" {
				Player = "O"
			} else {
				Player = "X"
			}
		}
	}
}
