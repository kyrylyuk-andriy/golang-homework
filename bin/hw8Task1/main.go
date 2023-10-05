package main

import (
	"context"
	"flag"
	"sync"

	"github.com/kyrylyuk-andriy/golang-homework/hw8Task1"
)

func main() {
	// number of orders
	a := flag.Int("count", 10, "How many orders we need to generate")
	flag.Parse()
	ctx := context.Background()
	ch := make(chan hw8Task1.CustomerOrder)

	var wg sync.WaitGroup
	// ctx := context.Background()

	wg.Add(2)
	go hw8Task1.GenerateCustomerOrder(ch, &wg, *a, ctx)
	go hw8Task1.ProcessOrders(ch, &wg, ctx)
	wg.Wait()
}
