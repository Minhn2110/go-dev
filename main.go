package main

import (
	"fmt"
)

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}
func sendOnly(c chan<- int) {
	c <- 2 // OK
	// error
	//fmt.Printf("Received: %d\n", <-c)
}

func recieveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
	// error
	//c <- 2
}

func main() {

	// Receive & Send
	//myChan := make(chan int)
	//go channel.ReceiveAndSend(myChan)
	//myChan <- 1
	//fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)

	// Receive Only
	//go recieveOnly(myChan)
	//myChan <- 1

	// Send Only
	//go sendOnly(myChan)
	//fmt.Printf("Value from sendOnly: %d\n", <-myChan)

	// Check is alive
	//channel.CheckChannelIsAlive()

}
