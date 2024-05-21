package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

//reflect.DeepEqual: useful for seeing if any two variables are the same.
//So while using reflect.DeepEqual is a convenient way of comparing slices (and other things),
//you must be careful when using it because reflect.DeepEqual is not "type safe" -
//the code will compile even if you did something a bit silly; like when we set want := "bob"
