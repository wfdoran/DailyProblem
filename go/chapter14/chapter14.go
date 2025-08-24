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
