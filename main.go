package main

import (
	"context"
	"fmt"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"log"
	"net/http"
	"time"
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

	//webhook.StartServer()

	//if err := run(context.Background()); err != nil {
	//	log.Fatal(err)
	//}

	now := time.Now().UTC()
	fmt.Println(now)
}

func run(ctx context.Context) error {
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("App URL", listener.URL())
	return http.Serve(listener, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello from ngrok-go!</h1>")
}
