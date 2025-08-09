package DailyPuzzle

import (
	"fmt"
	"reflect"
	"sort"
)

func AnagramIndices(w string, s string) []int {
	var rv []int

	w_map := make(map[rune]int)
	s_map := make(map[rune]int)

	for _, c := range w {
		w_map[c] += 1
	}
	// fmt.Println(w_map)

	ss := []rune(s)

	w_len := len(w)

	for i := range len(s) {
		s_map[ss[i]] += 1
		// fmt.Println(i, s_map)

		if i >= w_len-1 {
			lead_idx := i - w_len + 1
			if reflect.DeepEqual(s_map, w_map) {
				rv = append(rv, lead_idx)
			}

			remove_char := ss[lead_idx]
			s_map[remove_char] -= 1
			if s_map[remove_char] == 0 {
				delete(s_map, remove_char)
			}
		}
	}

	return rv
}

func StringReverse(s string) string {
	ss := []rune(s)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return string(ss)
}

func PalendromePairs(ss []string) [][2]int {
	var a []string
	var b []string
	idx := make(map[string]int)

	for i, s := range ss {
		a = append(a, s)
		b = append(b, StringReverse(s))
		idx[s] = i
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, s := range b {
		b[i] = StringReverse(s)
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(idx)

	var rv [][2]int

	return rv
}
