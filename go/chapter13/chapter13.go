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
