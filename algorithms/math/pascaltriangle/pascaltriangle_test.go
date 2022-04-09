package pascaltriangle

import (
	"reflect"
	"testing"
)

var cases = []struct {
	number int
	result []int
}{
	{number: 0, result: []int{1}},
	{number: 1, result: []int{1, 1}},
	{number: 2, result: []int{1, 2, 1}},
	{number: 3, result: []int{1, 3, 3, 1}},
	{number: 4, result: []int{1, 4, 6, 4, 1}},
	{number: 5, result: []int{1, 5, 10, 10, 5, 1}},
	{number: 6, result: []int{1, 6, 15, 20, 15, 6, 1}},
	{number: 7, result: []int{1, 7, 21, 35, 35, 21, 7, 1}},
}

func TestPascalTriangle(t *testing.T) {
	for _, input := range cases {
		result := pascalTriangle(input.number)

		if !reflect.DeepEqual(result, input.result) {
			t.Errorf("FAIL: number=%v, expected=%v, result=%v", input.number, input.result, result)
		}
	}
}
