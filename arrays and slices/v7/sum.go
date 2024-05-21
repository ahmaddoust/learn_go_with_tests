package main

import "fmt"

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in.
// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		tail := numbers[1:]
// 		sums = append(sums, Sum(tail))
// 	}

//		return sums
//	}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	// Test with an empty slice
	result := SumAllTails([]int{})
	fmt.Println(result) // Should output: [0]
}
