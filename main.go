package main

import (
	"sync"
	"time"

	"github.com/voltgizerz/go-concurrent/config"
)

func main() {
	start := time.Now()
	log := config.SetupLog()

	number := 0
	var wg sync.WaitGroup
	var mut sync.Mutex

	wg.Add(5)
	go func() {
		defer wg.Done()

		mut.Lock()
		for i := 0; i < 100000; i++ {
			number++
		}
		mut.Unlock()
	}()

	go func() {
		defer wg.Done()

		mut.Lock()
		for i := 0; i < 100000; i++ {
			number++
		}
		mut.Unlock()
	}()

	go func() {
		defer wg.Done()

		mut.Lock()
		for i := 0; i < 100000; i++ {
			number++
		}
		mut.Unlock()
	}()

	go func() {
		defer wg.Done()

		mut.Lock()
		for i := 0; i < 100000; i++ {
			number++
		}
		mut.Unlock()
	}()

	go func() {
		defer wg.Done()
		mut.Lock()
		for i := 0; i < 100000; i++ {
			number++
		}
		mut.Unlock()
	}()
	wg.Wait()

	log.Printf("Number printed %d, took %v\n", number, time.Since(start))
	log.Println("All process finished")
}
