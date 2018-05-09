package main

import (
	"testing"
)

func TestSum2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got int, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("Got %d. Want %d. Given %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers (array)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum2(numbers)
		want := 15
		assertCorrectMessage(t, got, want, numbers)
	})

	t.Run("collection of any size (slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum2(numbers)
		want := 6
		assertCorrectMessage(t, got, want, numbers)
	})
}
