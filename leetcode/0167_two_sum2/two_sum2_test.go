package twosum2

import (
	"reflect"
	"testing"
)

type arg struct {
	numbers []int
	target  int
}

func TestTwoSum2(t *testing.T) {
	testData := []arg{
		arg{
			numbers: []int{2, 7, 11, 15},
			target:  9,
		},
	}

	expectedData := [][]int{
		{1, 2},
	}

	for index, data := range testData {
		if res := twoSum2(data.numbers, data.target); !reflect.DeepEqual(res, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], res)
		}
	}
}
