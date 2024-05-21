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
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//In this implementation, we are worrying less about capacity. We start with an empty slice sums
//and append to it the result of Sum as we work through the varargs.

//the append function which takes a slice and a new value, then returns a
//new slice with all the items in it.
