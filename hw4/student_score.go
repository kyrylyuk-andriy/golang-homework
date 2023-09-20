package hw4

import (
	"fmt"
)

func Average_score() {
	var scores = []float32{4, 5.4, 3, 2.5, 6}
	var as, sum float32
	for _, v := range scores {
		sum = sum + v
	}
	as = sum / float32(len(scores))
	fmt.Printf("Average student score is %g \n", as)
}
