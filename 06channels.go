package interview

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateInt() {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			out <- rand.Intn(10)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()

		for item := range out {
			fmt.Println(item)
		}
	}()

	wg.Wait()

}
