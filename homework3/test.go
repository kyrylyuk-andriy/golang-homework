package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// i would make Steven life easier to gues between 10 and 20
	min, max := 10, 20
	secretNumber := rand.Intn(max-min) + min
	fmt.Printf("Enter number between %d and %d: \n", min, max)
	var guess int
	fmt.Scan(&guess)
	attempts := 0
	
	for {
		attempts++
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			fmt.Println("Correct, you Legend! You guessed right after")
			break
		}
	}
}