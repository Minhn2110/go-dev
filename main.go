package main

import (
	"fmt"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error retrieving hostname: %v\n", err)
		return
	}

	fmt.Printf("The hostname is: %s\n", host)
}
