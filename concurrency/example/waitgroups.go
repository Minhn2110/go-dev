package example

import (
	"fmt"
	"sync"
	"time"
)

func WorkerWcg(id int) {
	fmt.Println("ggg")

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func TestWcg() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			WorkerWcg(i)
		}()
	}

	wg.Wait()
	fmt.Println("Donme")
}
