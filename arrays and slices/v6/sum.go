package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

//SumAllTails: calculate the totals of the "tails" of each slice. The tail of a collection is all -
//items in the collection except the first one (the "head").

//In this implementation, we are worrying less about capacity. We start with an empty slice sums -
//and append to it the result of Sum as we work through the varargs.

//the append function which takes a slice and a new value, then returns a -
//new slice with all the items in it.

//Example:  for slices "[]int{1,2,3,4} and []int{2, 3,4,5}" the answer:

//For the slice []int{1, 2, 3, 4}:

//The tail (all elements except the first) includes the elements [2, 3, 4]. -
//The sum of the tail is 2 + 3 + 4 = 9.

// For the slice []int{2, 3, 4, 5}:

// The tail includes the elements [3, 4, 5].
// The sum of the tail is 3 + 4 + 5 = 12.

//result = [9, 12]
