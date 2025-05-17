package DailyPuzzle

import (
	"reflect"
	"testing"
)

func TestAnagramIndices(t *testing.T) {

	{
		w := "ab"
		s := "abxaba"
		expect := []int{0, 3, 4}
		out := AnagramIndices(w, s)
		if !reflect.DeepEqual(expect, out) {
			t.Error("out = ", out, "expect", expect)
		}
	}
}
