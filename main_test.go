package main

import "testing"

// Test function
func TestCalculate(t *testing.T) {
    if Calculate(2) != 4 {
        t.Error("Expected 2 + 2 to equal 4")
    }
}

// Define a bunch of edge cases for testing a function
func TestTableCalculate(t *testing.T) {
    // Define test table of inputs
    var tests = []struct {
        input int
        expected int

    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    // Iterate over all test case inputs, run each one individually
    for _, testCase := range tests {
        if output := Calculate(testCase.input); output != testCase.expected {
            t.Error("test failed:\n Input: {}\n Expected: {}\n Actual: {}", testCase.input, testCase.expected, output)
        }
    }
}
