package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

<<<<<<< HEAD
	if want != got {
=======
	if got != want {
>>>>>>> 5809362c59db6b29db1ea5623135729b26b39bc2
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
