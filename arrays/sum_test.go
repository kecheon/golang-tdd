package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of array numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		if sum != 6 {
			t.Errorf("expected %d but got %d", 6, sum)
		}
	})

	t.Run("sum of numbers in a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	})
}
