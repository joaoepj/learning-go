package lg_misc

import (
	"reflect"
	"testing"
)

var testSet = []struct{
	in []int32
	out []int32
}{
	{[]int32{1, 2, 3, 4, 5}, []int32{5, 4, 3, 2, 1}},
}

func TestReverse(t *testing.T) {
	for _, test := range testSet {
		t.Run("TestReverse", func(t *testing.T) {
			if res := Reverse(test.in, []int32{}); !reflect.DeepEqual(res, test.out) {
				t.Error("got, want: ", res, test.out)
			}

		})
	}
}