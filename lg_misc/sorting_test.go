package lg_misc

import (
	"reflect"
	"testing"
)

var TestSet = []struct {
	in  []int32
	out []int32
}{
	{[]int32{100, 100, 50, 40, 40, 20, 10, 5}, []int32{5, 10, 20, 40, 40, 50, 100, 100}},
	{[]int32{100, 250, 200, 50, 40, 40, 20, 10, 5}, []int32{5, 10, 20, 40, 40, 50, 100, 100}},
}

func TestSortingInit(t *testing.T) {
	t.Run("TestSortingInit", func(t *testing.T) {

	})
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range TestSet {
		t.Run("TestInsertionSort", func(t *testing.T) {
			res := InsertionSort(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Error("got, want: ", res, tt.out)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range TestSet {
		t.Run("TestSelectionSort", func(t *testing.T) {
			res := SelectionSort(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Error("got, want: ", res, tt.out)
			}
		})
	}
}
