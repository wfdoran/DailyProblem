package DailyPuzzle

import (
	"reflect"
	"testing"
)

func TestArrayProd(t *testing.T) {

	{
		in := []int{1, 2, 3, 4, 5}
		want := []int{120, 60, 40, 30, 24}
		out := ArrayProd(in)

		if !reflect.DeepEqual(out, want) {
			t.Error("ArrayProd", in, "=", out, "want", want)
		}
	}
}

func TestArrayProd2(t *testing.T) {

	{
		in := []int{1, 2, 3, 4, 5}
		want := []int{120, 60, 40, 30, 24}
		out := ArrayProd2(in)

		if !reflect.DeepEqual(out, want) {
			t.Error("ArrayProd", in, "=", out, "want", want)
		}
	}
}
