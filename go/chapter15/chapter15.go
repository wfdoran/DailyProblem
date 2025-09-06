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

func RadixSort(a []uint32) {
	n := len(a)
	b := make([]uint32, n)

	from := a
	to := b

	for shift := 0; shift <= 24; shift += 8 {
		var count [256]int

		for i := range n {
			idx := (from[i] >> shift) & 0xff
			count[idx]++
		}

		var offset [256]int
		offset[0] = 0

		for idx := 1; idx < 256; idx++ {
			offset[idx] = offset[idx-1] + count[idx-1]
		}

		for i := range n {
			idx := (from[i] >> shift) & 0xff
			to[offset[idx]] = from[i]
			offset[idx]++
		}

		to, from = from, to
	}
}
