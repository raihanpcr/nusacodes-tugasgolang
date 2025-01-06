package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestTugasConcurrency(t *testing.T) {
	
	calculator := &Calculator{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculator.Add(100, 10)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		calculator.Subtract(100, 10)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculator.Multiply(100, 10)
	}()

	wg.Wait()

	fmt.Println("Test Selesai")
}