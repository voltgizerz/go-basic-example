package main

import (
	"sync"
	"time"
)

// StartGoroutine - sample go routine process using wait group
func StartGoroutine() {
	start := time.Now()

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
