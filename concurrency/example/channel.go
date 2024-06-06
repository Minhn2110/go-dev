package example

import (
	"fmt"
	"time"
)

func SimpleChannel() {
	messages := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)

	//fmt.Println(runtime.NumCPU())

}
