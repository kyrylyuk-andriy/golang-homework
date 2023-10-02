package hw7task2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenerateNumbers(ch1 chan int, ch2 chan int, min int, max int, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {
		num := rand.Intn(max-min+1) + min
		fmt.Println(num)
		ch1 <- num
	}
	close(ch1)

	a := <-ch2
	b := <-ch2
	if a > b {
		fmt.Printf("the largest number %d\n", a)
		fmt.Printf("the smallest number %d\n", b)
	} else {
		fmt.Printf("the largest number %d\n", b)
		fmt.Printf("the smallest number %d\n", a)
	}
}

func BiggestSmallestNumber(ch1 chan int, ch2 chan int, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	largest := 0
	smallest := max
	// for n := range ch {
	// 	fmt.Println(n)
	// }
	for n := range ch1 {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}
	fmt.Printf("LOG: the largest number %d\n", largest)
	fmt.Printf("LOG: the smallest number %d\n", smallest)

	ch2 <- smallest
	ch2 <- largest

	close(ch2)
}
