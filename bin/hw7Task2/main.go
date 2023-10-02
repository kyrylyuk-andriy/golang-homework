package main

import (
	"sync"

	hw7task2 "github.com/kyrylyuk-andriy/golang-homework/hw7Task2"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	min := 1
	max := 1000
	amount := 5
	var wg sync.WaitGroup

	wg.Add(2)

	go hw7task2.GenerateNumbers(ch1, ch2, min, max, amount, &wg)
	go hw7task2.BiggestSmallestNumber(ch1, ch2, max, &wg)

	wg.Wait()
}
