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
		lo, hi := SortInterval2(in)

		if lo != want_lo || hi != want_hi {
			t.Error("SortInterval wrong")
		}

	}

	{
		in := []int{3, 7, 1, 8, 9}
		want_lo := 0
		want_hi := 2
		lo, hi := SortInterval2(in)

		if lo != want_lo || hi != want_hi {
			t.Error("SortInterval wrong")
		}

	}

}

func TestMaxArraySum(t *testing.T) {
	{
		in := []int{34, -50, 42, 14, -5, 86}
		want := 137
		actual := MaxArraySum(in)
		if actual != want {
			t.Errorf("actual %d, want %d", actual, want)
		}
	}

	{
		in := []int{-5, -1, -8, -9}
		want := 0
		actual := MaxArraySum(in)
		if actual != want {
			t.Errorf("actual %d, want %d", actual, want)
		}
	}
}

func TestSmallerRight(t *testing.T) {
	{
		in := []int{3, 4, 9, 6, 1}
		expect := []int{1, 1, 2, 1, 0}
		out := SmallerRight(in)
		if !reflect.DeepEqual(expect, out) {
			t.Error("ArrayProd", in, "=", out, "expect", expect)
		}
	}

	{
		in := []int{1, 1, 1, 1, 1}
		expect := []int{0, 0, 0, 0, 0}
		out := SmallerRight(in)
		if !reflect.DeepEqual(expect, out) {
			t.Error("ArrayProd", in, "=", out, "expect", expect)
		}
	}
}
