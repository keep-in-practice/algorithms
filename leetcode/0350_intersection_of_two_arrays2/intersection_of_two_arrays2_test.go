package intersectionof2arrays2

import (
	"reflect"
	"testing"
)

func TestInteract(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	expectedData := []int{2, 2}

	if res := intersect(nums1, nums2); !reflect.DeepEqual(res, expectedData) {
		t.Errorf("expected %v, got %v", expectedData, res)
	}
}
