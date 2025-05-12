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

func TestArrayProd3(t *testing.T) {

	{
		in := []int{1, 2, 3, 4, 5}
		want := []int{120, 60, 40, 30, 24}
		out := ArrayProd3(in)

		if !reflect.DeepEqual(out, want) {
			t.Error("ArrayProd", in, "=", out, "want", want)
		}
	}
}

func TestSortInterval(t *testing.T) {

	{
		in := []int{3, 7, 5, 6, 9}
		want_lo := 1
		want_hi := 3
		lo, hi := SortInterval(in)

		if lo != want_lo || hi != want_hi {
			t.Error("SortInterval wrong")
		}

	}

	{
		in := []int{3, 7, 1, 8, 9}
		want_lo := 0
		want_hi := 2
		lo, hi := SortInterval(in)

		if lo != want_lo || hi != want_hi {
			t.Error("SortInterval wrong")
		}

	}

}
