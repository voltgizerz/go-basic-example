package main

import (
	"sync"
	"testing"
	"time"
)

func TestStartGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex

	tests := []struct {
		name    string
		numGoroutines int
	}{
		{
			name: "Two goroutines",
			numGoroutines: 2,
		},
		{
			name: "Five goroutines",
			numGoroutines: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			number := 0

			wg.Add(tt.numGoroutines)
			for i := 0; i < tt.numGoroutines; i++ {
				go func() {
					defer wg.Done()

					mut.Lock()
					for i := 0; i < 100000; i++ {
						number++
					}
					mut.Unlock()
				}()
			}
			wg.Wait()

			expected := tt.numGoroutines * 100000
			if number != expected {
				t.Errorf("StartGoroutine() = %d, want %d", number, expected)
			}

			elapsed := time.Since(start)
			if elapsed > 2*time.Second {
				t.Errorf("StartGoroutine() took too long: %v", elapsed)
			}
		})
	}
}