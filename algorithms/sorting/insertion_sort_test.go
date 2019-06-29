package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{23, 1, 53, 543, 98, 4}

	got := InsertionSort(nums)
	want := []int{1, 4, 23, 53, 98, 543}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
	}
}
