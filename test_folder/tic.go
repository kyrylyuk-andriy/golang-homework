package main

import (
	"fmt"
	"strconv"
)

var board = make(map[int]string)
var currentPlayer string
var gameActive bool

var winningCombinations = [][3]int{
	{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, // Rows
	{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, // Columns
	{1, 5, 9}, {3, 5, 7}, // Diagonals
}

func initializeBoard() {
	for i := 1; i <= 9; i++ {
		board[i] = " "
	}
}

func printBoard() {
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---|---|---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---|---|---")
	fmt.Println(" 7 | 8 | 9 ")
}

func checkWin() bool {
	for _, combo := range winningCombinations {
		if board[combo[0]] == currentPlayer && board[combo[1]] == currentPlayer && board[combo[2]] == currentPlayer {
			return true
		}
	}

	return false
}

func isBoardFull() bool {
	for i := 1; i <= 9; i++ {
		if board[i] == " " {
			return false
		}
	}
	return true
}

func main() {
	initializeBoard()
	currentPlayer = "X"
	gameActive = true

	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("Use numbers 1-9 to make your move.")
	printBoard()

	for gameActive {
		fmt.Printf("Player %s's turn.\n", currentPlayer)
		fmt.Print("Enter your move: ")

		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 9.")
			continue
		}

		position, err := strconv.Atoi(input)

		if err != nil || position < 1 || position > 9 || board[position] != " " {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		board[position] = currentPlayer
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
