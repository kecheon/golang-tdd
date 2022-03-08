package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing Counter 3 times", func(t *testing.T) {
		counter := Counter{}
		// int default value sets to 0
		fmt.Println(counter.value)
		counter.Inc()
		counter.Inc()
		counter.Inc()
		got := counter.Value()
		want := 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("run in multiple threads concurrent environment", func(t *testing.T) {
		counter := NewCounter()
		num := 1000
		var wg sync.WaitGroup
		wg.Add(num)
		for i := 0; i < num; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, 1000)
	})
}

func assertCounter(t testing.TB, got *Counter, expected int) {
	t.Helper()
	if got.Value() != expected {
		t.Errorf("got %d, want %d", expected, got.Value())
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
