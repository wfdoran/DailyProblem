package chapter12

func TowerOfHanoi(n int, src int, dst int) chan [2]int {
	ch := make(chan [2]int)
	go func(ch chan [2]int) {
		if n == 1 {
			x := [2]int{src, dst}
			ch <- x
		} else {
			spare := 6 - src - dst
			for x := range TowerOfHanoi(n-1, src, spare) {
				ch <- x
			}
			x := [2]int{src, dst}
			ch <- x
			for x := range TowerOfHanoi(n-1, spare, dst) {
				ch <- x
			}
		}
		close(ch)
	}(ch)

	return ch
}

func RegEx(r string, s string) bool {
	r_len := len(r)
	s_len := len(s)

	if r_len == 0 {
		return s_len == 0
	}

	r_head := r[0]
	r_tail := r[1:]

	if r_head == '*' {
		for i := range s_len + 1 {
			s_tail := s[i:]
			if RegEx(r_tail, s_tail) {
				return true
			}
		}
	}

	if s_len == 0 {
		return false
	}

	s_head := s[0]
	s_tail := s[1:]

	if r_head != '.' && r_head != s_head {
		return false
	}
	return RegEx(r_tail, s_tail)
}

func ArrayMinMax(a []int) (int, int) {
	n := len(a)

	if n == 1 {
		return a[0], a[0]
	}

	if n == 2 {
		if a[0] < a[1] {
			return a[0], a[1]
		} else {
			return a[1], a[0]
		}
	}

	m := n / 2
	min1, max1 := ArrayMinMax(a[:m])
	min2, max2 := ArrayMinMax(a[m:])

	var min3, max3 int
	if min1 < min2 {
		min3 = min1
	} else {
		min3 = min2
	}

	if max1 > max2 {
		max3 = max1
	} else {
		max3 = max2
	}
	return min3, max3
}
