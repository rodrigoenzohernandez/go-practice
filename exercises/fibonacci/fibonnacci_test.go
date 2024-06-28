package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	f := fibonacci()

	// Test the first 10 Fibonacci numbers
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i := 0; i < 10; i++ {
		result := f()
		assert.Equal(t, expected[i], result, "Fibonacci number at index %d should match expected value", i)
	}

	// Skipping to the 40th number
	for i := 10; i < 40; i++ {
		f() // Call f() to advance but ignore the result
	}
	result := f() // This should now be the 40th number
	assert.Equal(t, 102334155, result, "Fibonacci number at index 40 should match expected value")

	// Additional tests can be added here in a similar manner
}
