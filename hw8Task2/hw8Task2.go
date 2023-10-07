package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Question string
	Answer   int
}

func GenerateNumberAndSendQuestion(wg *sync.WaitGroup, ctx context.Context, ch chan string, referre chan Game) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		// fmt.Fprint(os.Stderr, "request cancelled\n")
		return
	case <-time.After(1 * time.Second):
		question := " What's your answer? Quess number between 1 and 10:"
		fmt.Println(question)
		answer := rand.Intn(10)
		fmt.Println(answer)
		game := Game{Question: question, Answer: answer}
		referre <- game
		ch <- question
		close(referre)
		close(ch)
	}
}

func Player(wg *sync.WaitGroup, ctx context.Context, ch chan string, result chan int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		// fmt.Fprint(os.Stderr, "request cancelled\n")
		return
	case question := <-ch:
		fmt.Printf("Player received: %s\n", question)
		answer := rand.Intn(10)
		fmt.Printf("Player's answer: %d", answer)
		result <- answer
		close(result)
	}
}

func Referee(wg *sync.WaitGroup, ctx context.Context, referre chan Game, result chan int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		// fmt.Fprint(os.Stderr, "request cancelled\n")
		return
	case answer := <-result:
		fmt.Printf("Referee received answer from Player: %d\n", answer)
		game := <-referre
		// fmt.Printf("Referee received question from Game center : %s\n", game)
		if game.Answer == answer {
			fmt.Println("Player won")
		} else {
			fmt.Println("Player lost")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result := make(chan int)
	referee := make(chan Game)
	answer := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go GenerateNumberAndSendQuestion(&wg, ctx, answer, referee)
	go Player(&wg, ctx, answer, result)
	go Referee(&wg, ctx, referee, result)

	wg.Wait()

	// Cancel the game
	cancel()
}
