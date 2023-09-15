package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	min, max := 10, 20
	secretNumber := rand.Intn(max-min) + min

	fmt.Println("Guess a number between 10 and 20")
	fmt.Println("Please input your guess")

	attempts := 0
	for {
		attempts++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error reading input. Please try again", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		fmt.Println("Your guess is", guess)

		if guess != secretNumber {
			fmt.Println("Your guess not right. Try again")
		} else {
			fmt.Println("Correct, safe is open")
			break
		}
	}
}