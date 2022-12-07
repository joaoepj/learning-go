package lg_misc

import (
	//"reflect"
	"testing"
)

func TestCuriosity(t *testing.T) {
		t.Run("TestCuriosity", func(t *testing.T) {
			Curiosity()
			if 2 != 1 {
				t.Error("got, want: ", 1)
			}
		})
}