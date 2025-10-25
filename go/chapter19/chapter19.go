package chapter19

func strHash(s string) int {
	rv := 0
	for _, ch := range s {
		rv += int(ch)
	}
	return rv
}

func RabinKarp(target string, s string) []int {
	target_hash := strHash(target)
	n := len(target)
	m := len(s)

	rv := []int{}

	if m < n {
		return rv
	}

	curr_hash := 0
	for i, ch := range s {
		curr_hash += int(ch)
		if i < n-1 {
			continue
		}

		if curr_hash == target_hash {
			if target == s[i-n+1:i+1] {
				rv = append(rv, i-n+1)
			}
		}

		curr_hash -= int([]rune(s)[i-n+1])
	}

	return rv

}
