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

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(ArrayProd(x))
}
