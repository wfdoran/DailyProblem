package chapter13

var hist = map[int64]int64{}

func NumClimbs(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	val, found := hist[n]
	if found {
		return val
	}

	out := NumClimbs(n-1) + NumClimbs(n-2)
	hist[n] = out
	return out
}

func DecodeCount(s string) int {
	n := len(s)
	count := make([]int, n+1)

	count[n] = 1
	if s[n-1] == '0' {
		count[n-1] = 0
	} else {
		count[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			count[i] = 0
			continue
		}

		count[i] = count[i+1]

		if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			count[i] += count[i+2]
		}
	}
	return count[0]
}
