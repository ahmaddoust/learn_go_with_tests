package main

import "testing"

// We will now use the slice type which allows us to have collections of any size.
// The syntax is very similar to arrays, you just omit the size when declaring them
// mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
