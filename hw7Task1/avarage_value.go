package hw7Task1

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumbers(randomNumbers chan int, randomAmount int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < randomAmount; i++ {
		num := rand.Intn(10000)
		randomNumbers <- num
		fmt.Println(num)
	}
	close(randomNumbers)
}

func CalculateAvarageNumber(randomNumbers chan int, avarageResult chan float32, randomAmount int) {
	sum := 0
	for v := range randomNumbers {
		sum = v + sum
	}
	avarage := float32(sum) / float32(randomAmount)
	avarageResult <- avarage
	close(avarageResult)
}
