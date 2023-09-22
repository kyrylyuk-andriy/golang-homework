package main

import (
	"fmt"
	"strconv"
)

var board = make(map[string]string)
var currentPlayer string
var gameActive bool

func initializeBoard() {
	for i := 1; i <= 9; i++ {
		board[strconv.Itoa(i)] = " "
	}
}

func printBoard() {
	fmt.Println("  1 2 3")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s|%s|%s\n", i+1, board[strconv.Itoa(i*3+1)], board[strconv.Itoa(i*3+2)], board[strconv.Itoa(i*3+3)])
		if i < 2 {
			fmt.Println("  -----")
		}
	}
	fmt.Println()
}

func checkWin() bool {
	winPatterns := [8][3]string{
		{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"},
		{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"},
		{"1", "5", "9"}, {"3", "5", "7"},
	}

	for _, pattern := range winPatterns {
		if board[pattern[0]] == currentPlayer && board[pattern[1]] == currentPlayer && board[pattern[2]] == currentPlayer {
			return true
		}
	}

	return false
}

func isBoardFull() bool {
	for i := 1; i <= 9; i++ {
		if board[strconv.Itoa(i)] == " " {
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
	printBoard()

	for gameActive {
		fmt.Printf("Player %s's turn.\n", currentPlayer)
		fmt.Print("Enter cell number (1-9): ")

		var cellNum int
		_, err := fmt.Scanf("%d", &cellNum)

		if err != nil || cellNum < 1 || cellNum > 9 || board[strconv.Itoa(cellNum)] != " " {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		board[strconv.Itoa(cellNum)] = currentPlayer
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
