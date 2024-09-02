package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		rows, cols int
		expected   [][]int
	}
	tests := []testCase{
		{3, 3, [][]int{
			{0, 0, 0},
			{0, 1, 2},
			{0, 2, 4},
		}},
		{4, 4, [][]int{
			{0, 0, 0, 0},
			{0, 1, 2, 3},
			{0, 2, 4, 6},
			{0, 3, 6, 9},
		}},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{5, 7, [][]int{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 1, 2, 3, 4, 5, 6},
				{0, 2, 4, 6, 8, 10, 12},
				{0, 3, 6, 9, 12, 15, 18},
				{0, 4, 8, 12, 16, 20, 24},
			}},
			{0, 0, [][]int{}},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := createMatrix(test.rows, test.cols)
		if !reflect.DeepEqual(output, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed: %v x %v matrix
Expecting:
%v
Actual:
%v
Fail
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed: %v x %v matrix
Expecting:
%v
Actual:
%v
Pass
`, test.rows, test.cols, formatMatrix(test.expected), formatMatrix(output))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true

func formatMatrix(matrix [][]int) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%v\n", row)
	}
	return result
}
