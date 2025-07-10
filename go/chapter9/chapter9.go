package chapter9

import "cmp"

func HeapDown[T cmp.Ordered](a []T, start_idx int) {
	n := len(a)

	idx := start_idx

	for {
		swap_idx := idx

		left := 2*idx + 1
		right := 2*idx + 2

		if right < n && a[idx] > a[right] {
			swap_idx = right
		}

		if left < n && a[idx] > a[left] {
			swap_idx = left
		}

		if swap_idx == idx {
			break
		}

		a[idx], a[swap_idx] = a[swap_idx], a[idx]
		idx = swap_idx
	}
}

func HeapUp[T cmp.Ordered](a []T, start_idx int) {
	idx := start_idx

	for idx > 0 {
		parent := (idx - 1) / 2

		if a[parent] <= a[idx] {
			break
		}

		a[idx], a[parent] = a[parent], a[idx]
		idx = parent
	}
}

func HeapInsert[T cmp.Ordered](a []T, val T) []T {
	a = append(a, val)
	HeapUp(a, len(a)-1)
	return a
}

func HeapDelete[T cmp.Ordered](a []T) (T, []T) {
	rv := a[0]
	n := len(a)
	a[0] = a[n-1]
	a = a[:n-1]
	HeapDown(a, 0)
	return rv, a
}
