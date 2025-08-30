package chapter15

func DutchFlagSort(a []rune) {
	r_pos := 0
	b_pos := len(a) - 1

	for i := 0; i <= b_pos; i++ {
		if a[i] == 'R' {
			a[i], a[r_pos] = a[r_pos], a[i]
			r_pos++
		} else if a[i] == 'B' {
			a[i], a[b_pos] = a[b_pos], a[i]
			b_pos--
			i--
		}
	}
}
