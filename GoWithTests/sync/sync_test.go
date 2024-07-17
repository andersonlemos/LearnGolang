package _sync

import (
	"sync"
	"testing"
)

func checkCounter(t *testing.T, c *Counter, expected int) {
	t.Helper()
	if c.Value() != expected {
		t.Errorf("expected %d, got %d", expected, c.Value())
	}
}

func TestCounter(t *testing.T) {
	t.Run("increase the counter 3 times results in 3 value", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		checkCounter(t, counter, 3)
	})
	t.Run("runs concurrently safely", func(t *testing.T) {
		expectedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(expectedCounter)

		for i := 0; i < expectedCounter; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Increment()
				wg.Done()
			}(&wg)
		}

		wg.Wait()
		checkCounter(t, counter, expectedCounter)
	})
}
