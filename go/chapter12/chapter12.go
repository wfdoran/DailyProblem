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
