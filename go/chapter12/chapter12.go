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
