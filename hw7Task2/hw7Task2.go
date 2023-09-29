package hw7task2

import (
	"fmt"
	"math/rand"
	"sync"
)

func GenerateNumbers(ch chan []int, min int, max int, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	var numbers = []int{}
	for i := 0; i < amount; i++ {
		num := rand.Intn(max-min) + min
		fmt.Println(num)
		numbers = append(numbers, num)
	}
	ch <- numbers
}

func BiggestSmallestNumber(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	largest := 0
	smallest := 0
	var numbersToChannel = []int{}

	for _, n := range <-ch {
		if n > largest {
			largest = n
			fmt.Println(largest)
		}
		if n < smallest {
			smallest = n
			fmt.Println(smallest)
		}
	}
	numbersToChannel = append(numbersToChannel, smallest)
	numbersToChannel = append(numbersToChannel, largest)
	ch <- numbersToChannel
}

func ReadFromChannel(ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var numbers = []int{}
	fmt.Println("Reading from channel")
	for _, n := range <-ch {
		numbers = append(numbers, n)
	}

	smallest := numbers[0]
	largest := numbers[1]
	fmt.Printf("the smallest number is %d\n", smallest)
	fmt.Printf("the largest number is %d\n", largest)
}
