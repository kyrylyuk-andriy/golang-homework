package hw8Task1

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

type CustomerOrder struct {
	ProductName  string
	CustomerName string
	Quantity     int
	Country      string
	City         string
	Price        int
}

func GenerateRandomStringData(name string, number int) []string {
	stringSlice := make([]string, number)
	for i := 0; i < number; i++ {
		newString := fmt.Sprintf("%s%d", name, i)
		stringSlice[i] = newString
	}
	return stringSlice
}

func GenerateCustomerOrder(ch chan CustomerOrder, wg *sync.WaitGroup, orderNumbers int, ctx context.Context) {
	defer wg.Done()
	var order CustomerOrder
	ProductNames := GenerateRandomStringData("Product", orderNumbers)
	CustomerNames := GenerateRandomStringData("Customer", orderNumbers)
	CountryNames := GenerateRandomStringData("Country", orderNumbers)
	CityNames := GenerateRandomStringData("City", orderNumbers)

	fmt.Println(ProductNames)
	fmt.Println(CustomerNames)
	fmt.Println(CountryNames)
	fmt.Println(CityNames)

	for i := 0; i < orderNumbers; i++ {
		order = CustomerOrder{
			ProductName:  ProductNames[rand.Intn(len(ProductNames))],
			CustomerName: CustomerNames[rand.Intn(len(CustomerNames))],
			Quantity:     rand.Intn(orderNumbers) + 1,
			Country:      CountryNames[rand.Intn(len(CountryNames))],
			City:         CityNames[rand.Intn(len(CityNames))],
			Price:        rand.Intn(100),
		}
		ch <- order
		fmt.Println(order)
	}
	close(ch)
}

func ProcessOrders(ch chan CustomerOrder, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	totalOrderPrice := 0

	for eachOrder := range ch {
		OrderPrice := eachOrder.Price * eachOrder.Quantity
		fmt.Printf("Order from: %s, product name: %s, price %d, Quantity: %d \n", eachOrder.CustomerName, eachOrder.ProductName, eachOrder.Price, eachOrder.Quantity)
		totalOrderPrice = totalOrderPrice + OrderPrice
	}
	fmt.Printf("Total price for all orders is %d\n", totalOrderPrice)
}
