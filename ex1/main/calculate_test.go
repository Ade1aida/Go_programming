package main

import "testing"

func TestCalculate(t *testing.T) {

	testCases := []struct {
		expression []string
		expected   float64
	}{
		{[]string{"2", "+", "3"}, 5},
		{[]string{"2", "-", "3"}, -1},
		{[]string{"2", "*", "3"}, 6},
		{[]string{"6", "/", "3"}, 2},
	}

	for _, testCase := range testCases {
		result := calculate(testCase.expression)
		if result != testCase.expected {
			t.Errorf("Expression '%s' returned incorrect result: expected %f, got %f", testCase.expression, testCase.expected, result)
		}
	}
}
