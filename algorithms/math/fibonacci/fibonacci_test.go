package fibonacci

import (
	"reflect"
	"testing"
)

var cases = []struct {
	number   int
	expected []int
}{
	{
		number:   1,
		expected: []int{1},
	},
	{
		number:   2,
		expected: []int{1, 1},
	},
	{
		number:   3,
		expected: []int{1, 1, 2},
	},
	{
		number:   4,
		expected: []int{1, 1, 2, 3},
	},
	{
		number:   5,
		expected: []int{1, 1, 2, 3, 5},
	},
	{
		number:   6,
		expected: []int{1, 1, 2, 3, 5, 8},
	},
	{
		number:   7,
		expected: []int{1, 1, 2, 3, 5, 8, 13},
	},
	{
		number:   8,
		expected: []int{1, 1, 2, 3, 5, 8, 13, 21},
	},
	{
		number:   9,
		expected: []int{1, 1, 2, 3, 5, 8, 13, 21, 34},
	},
	{
		number:   10,
		expected: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
	},
}

func TestFibonacci(t *testing.T) {
	for _, input := range cases {
		result := fibonacci(input.number)

		if !reflect.DeepEqual(result, input.expected) {
			t.Errorf("FAIL: number=%v, expected=%v, result=%v", input.number, input.expected, result)
		}
	}
}
