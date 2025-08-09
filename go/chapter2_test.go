package DailyPuzzle

import (
	"fmt"
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

func TestPalendromePairs(t *testing.T) {
	{
		in := []string{"foo", "fooab", "foobb"}
		out := PalendromePairs(in)

		fmt.Println(out)
	}
}
