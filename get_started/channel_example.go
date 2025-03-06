package main

import (
	"fmt"
	"time"
)

type order struct {
	ID     int
	Status string
}

func processOrder(orderID int, StatusChannel chan order) {
	time.Sleep(time.Second * 2)
	StatusChannel <- order{ID: orderID, Status: "Processed"}

}
func main() {
	StatusChannel := make(chan order)

	for i := 1; i <= 5; i++ {
		go processOrder(i, StatusChannel)
	}

	for i := 1; i <= 5; i++ {
		processedOrder := <-StatusChannel
		fmt.Println(processedOrder.ID, processedOrder.Status)
	}
}
