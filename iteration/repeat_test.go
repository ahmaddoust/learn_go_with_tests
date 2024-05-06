package main

import (
	"fmt"
	"strings"
	"testing"
)

// Repeat concatenates the specified `character` string `repeatCount` times.
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// BenchmarkRepeat benchmarks the Repeat function for a specified repeatCount.
func BenchmarkRepeat(b *testing.B) {
	repeatCount := 1000 // Specify the desired repeatCount

	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

// ExampleRepeat demonstrates how to use the Repeat function.
func ExampleRepeat() {
	result := Repeat("abc", 3)
	fmt.Println(result)
	// Output: abcabcabc
}

// Explanation
// Updated Repeat Function:
// The Repeat function now accepts two parameters: character (the string to repeat) and repeatCount (the number of times to repeat the string).
// Inside the function, repeatCount determines the loop iteration count to build the repeated string using a strings.Builder.
// Updated BenchmarkRepeat Function:
// The BenchmarkRepeat function now defines repeatCount (e.g., 1000) to specify how many times the character should be repeated during benchmarking.
// The benchmark loop calls Repeat("a", repeatCount) to benchmark the Repeat function with the specified repeatCount.
// Example Function (ExampleRepeat):
// The ExampleRepeat function demonstrates how to use the Repeat function by calling it with "abc" and 3 as arguments (Repeat("abc", 3)) and then printing the result using fmt.Println.
// The comment // Output: abcabcabc specifies the expected output when running ExampleRepeat. This comment is recognized by the go doc tool to generate example documentation.
// Running Example and Documentation
// To run the ExampleRepeat function and view the documentation, you can use the go test and go doc commands:

// go test -run Example
// This command runs the Example functions (including ExampleRepeat) and displays their output. It's useful for verifying that the example functions work as expected.

// To view the documentation for the Repeat function, you can use go doc:

// go doc -cmd Repeat
// This command displays the documentation for the Repeat function, including the function signature and the example usage provided in ExampleRepeat.

// By following these steps, you'll have a well-documented Repeat function with example usage, along with a benchmark function that allows flexibility in specifying the number of repetitions for benchmarking purposes.
