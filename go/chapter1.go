package DailyPuzzle

import (
	"fmt"
)

func ArrayProd(a []int) []int {
	n := len(a)
	rv := make([]int, n)

	prod := 1
	for _, v := range a {
		prod *= v
	}

	for i, v := range a {
		rv[i] = prod / v
	}
	return rv
}

func ArrayProdWorker(a []int) ([]int, int) {
	n := len(a)
	rv := make([]int, n)

	if n == 1 {
		rv[0] = 1
		return rv, a[0]
	}

	k := n / 2

	sub1, prod1 := ArrayProdWorker(a[:k])
	sub2, prod2 := ArrayProdWorker(a[k:])

	for i, v := range sub1 {
		rv[i] = v * prod2
	}
	for i, v := range sub2 {
		rv[i+k] = v * prod1
	}
	return rv, prod1 * prod2
}

func ArrayProd2(a []int) []int {
	rv, _ := ArrayProdWorker(a)
	return rv
}

func ArrayProd3(a []int) []int {
	n := len(a)
	rv := make([]int, n)
	rv[0] = 1
	for i := range n - 1 {
		rv[i+1] = rv[i] * a[i]
	}

	trail := 1
	for i := n - 1; i >= 0; i-- {
		rv[i] *= trail
		trail *= a[i]
	}
	return rv
}

func SortInterval(a []int) (int, int) {
	n := len(a)
	var stack []int

	sort_lo := n
	for i := 0; i < n; i++ {
		v := a[i]
		for len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			if len(stack) < sort_lo {
				sort_lo = len(stack)
			}
		}
		stack = append(stack, v)
	}

	stack = stack[:0]
	sort_hi := -1
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for len(stack) > 0 && v > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			x := n - 1 - len(stack)
			if x > sort_hi {
				sort_hi = x
			}
		}
		stack = append(stack, v)

	}
	fmt.Println(sort_lo, sort_hi)
	return sort_lo, sort_hi
}

func SortInterval2(a []int) (int, int) {
	n := len(a)

	max_seen := a[0]
	sort_hi := -1
	for i := 0; i < n; i++ {
		v := a[i]
		if v > max_seen {
			max_seen = v
		}
		if v < max_seen {
			sort_hi = i
		}
	}

	min_seen := a[n-1]
	sort_lo := n
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if v < min_seen {
			min_seen = v
		}
		if v > min_seen {
			sort_lo = i
		}
	}
	return sort_lo, sort_hi
}

func max(a ...int) int {
	rv := a[0]

	for _, v := range a[1:] {
		if v > rv {
			rv = v
		}
	}
	return rv
}

func MaxArraySum(a []int) int {
	rv := 0
	curr := 0

	for _, v := range a {
		curr = max(curr+v, v, 0)
		rv = max(curr, rv)
	}
	return rv
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(ArrayProd(x))
}
