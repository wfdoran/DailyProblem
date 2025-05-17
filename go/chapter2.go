package DailyPuzzle

import (
	"reflect"
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
