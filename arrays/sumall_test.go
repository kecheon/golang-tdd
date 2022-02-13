package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2}
	numbers2 := []int{3, 4}
	got := SumAll(numbers1, numbers2)
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum all tails", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 3, 4}
		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 7}
		checkSums(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{}
		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 0}
		checkSums(t, got, want)
	})
}
