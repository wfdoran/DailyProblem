package chapter11

func FenwickUpdate(a []int, idx int, amt int) {
	n := len(a)

	for ; idx < n; idx += idx & -idx {
		a[idx] += amt
	}
}

func FenwickSum(a []int, idx int) int {
	rv := 0
	for ; idx > 0; idx -= idx & -idx {
		rv += a[idx]
	}
	return rv
}

func FenwickRange(a []int, left int, right int) int {
	return FenwickSum(a, right) - FenwickSum(a, left-1)
}
