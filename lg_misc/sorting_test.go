package lg_misc

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		in  []int32
		out []int32
	}{
		{[]int32{100, 100, 50, 40, 40, 20, 10, 5}, []int32{5, 10, 20, 40, 40, 50, 100, 100}},
	}

	for _, tt := range tests {
		t.Run("Test1", func(t *testing.T) {
			res := InsertionSort(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Error("got, want: ", res, tt.out)
			}
		})
	}
}
