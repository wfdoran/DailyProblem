package chapter14

type flight struct {
	src string
	dst string
}

func FlightItinerary(f []flight, start string) []string {
	n := len(f)

	depth := 0
	state := make([]int, n)

	state[depth] = 0

	for {
		if state[depth] >= n {
			if depth == 0 {
				break
			}
			depth--
			state[depth]++
		}

		good := true

		for i := range depth {
			if state[i] == state[depth] {
				good = false
			}
		}

		if depth == 0 && f[state[depth]].src != start {
			good = false
		}
		if depth > 0 && f[state[depth]].src != f[state[depth-1]].dst {
			good = false
		}

		if !good {
			state[depth] += 1
			continue
		}

		depth++

		if depth == n {
			rv := make([]string, n+1)
			rv[0] = start
			for i := range n {
				rv[i+1] = f[state[i]].dst
			}
			return rv
		}

		state[depth] = 0

	}
	return nil
}

func AndroidCount(n int) int {
	depth := 0
	state := make([]int, n)
	count := 0

	state[depth] = 1

	for {
		if state[depth] == 10 {
			if depth == 0 {
				return count
			}
			depth--
			state[depth]++
			continue
		}

		good := true

		for i := range depth {
			if state[i] == state[depth] {
				good = false
			}
		}

		if depth >= 1 {
			all_mask := 0
			for i := range depth + 1 {
				all_mask |= 1 << state[i]
			}
			back2_mask := (1 << state[depth]) | (1 << state[depth-1])

			jump_tests := [][3]int{{1, 3, 2}, {4, 6, 5}, {7, 9, 8}, {1, 7, 4}, {2, 8, 5}, {3, 9, 6}}

			for _, t := range jump_tests {
				if back2_mask == (1<<t[0])|(1<<t[1]) {
					if (all_mask>>t[2])&1 == 0 {
						good = false
					}
				}
			}
		}

		if !good {
			state[depth]++
			continue
		}

		if depth == n-1 {
			count++
			state[depth]++
			continue
		}

		depth++
		state[depth] = 1
	}

}
