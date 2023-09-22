package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string
var gameActive bool

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("  1 2 3")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s|%s|%s\n", i+1, board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("  -----")
		}
	}
	fmt.Println()
}

func checkWin() bool {
	for i := 0; i < 3; i++ {
		// Check rows and columns
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}

	return false
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func main() {
	initializeBoard()
	currentPlayer = "X"
	gameActive = true

	fmt.Println("Welcome to Tic-Tac-Toe!")
	printBoard()

	for gameActive {
		fmt.Printf("Player %s's turn.\n", currentPlayer)
		fmt.Print("Enter row (1-3) and column (1-3) separated by space: ")

		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)

		if err != nil || row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != " " {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		board[row-1][col-1] = currentPlayer
		printBoard()

		if checkWin() {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			gameActive = false
		} else if isBoardFull() {
			fmt.Println("It's a tie!")
			gameActive = false
		} else {
			// Switch player
			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		}
	}
}
