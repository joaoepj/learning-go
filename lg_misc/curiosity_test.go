package lg_misc

import (
	//"reflect"
	"testing"
)

func TestCuriosity(t *testing.T) {
		t.Run("Curiosity=1", func(t *testing.T) {
			Curiosity()
			if 1 != 1 {
				t.Error("got, want: ", 1)
			}
		})
                t.Run("Curiosity=2", func(t *testing.T) {
                        Curiosity()
                        if 2 != 1 {
                                t.Error("got, want: ", 1)
                        }
                })
}
