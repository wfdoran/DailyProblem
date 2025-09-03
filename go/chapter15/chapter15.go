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

func reverse(a []int, i, j int) {
	ii := i
	jj := j

	for ii < jj {
		a[ii], a[jj] = a[jj], a[ii]
		ii++
		jj--
	}
}

func PancakeSort(a []int) {
	n := len(a)

	for i := range n - 1 {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min_idx] {
				min_idx = j
			}
		}

		reverse(a, i, min_idx)
	}
}
