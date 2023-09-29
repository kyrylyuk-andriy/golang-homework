package main

import (
	"sync"

	hw7task2 "github.com/kyrylyuk-andriy/golang-homework/hw7Task2"
)

func main() {
	ch := make(chan []int)
	min := 1
	max := 1000
	amount := 5
	var wg sync.WaitGroup

	wg.Add(3)

	go hw7task2.GenerateNumbers(ch, min, max, amount, &wg)
	go hw7task2.BiggestSmallestNumber(ch, &wg)
	go hw7task2.ReadFromChannel(ch, &wg)

	wg.Wait()

}
