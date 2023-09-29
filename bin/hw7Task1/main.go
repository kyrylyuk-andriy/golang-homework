package main

import (
	"fmt"
	"time"

	"github.com/kyrylyuk-andriy/golang-homework/hw7Task1"
)

func main() {
	randomNumbers := make(chan int)
	avarageResult := make(chan float32)
	randomNumber := 20

	go hw7Task1.RandomNumbers(randomNumbers, randomNumber)
	go hw7Task1.CalculateAvarageNumber(randomNumbers, avarageResult, randomNumber)
	n := <-avarageResult

	go fmt.Printf("avarage number is %v \n", n)
	time.Sleep(1 * time.Second)
}
