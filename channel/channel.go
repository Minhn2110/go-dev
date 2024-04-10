package channel

import "fmt"

func CheckChannelIsAlive() {
	myChan := make(chan int)

	go func() {
		myChan <- 3
	}()

	value, isAlive := <-myChan
	close(myChan) // Close

	fmt.Printf("Value from sendOnly: %d\n", value)
	fmt.Println("isAlive", isAlive)
}

func ReceiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}
