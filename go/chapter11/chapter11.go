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

type UnionFind struct {
	num_classes int
	class       []int
	class_size  []int
}

func UnionFindInit(num_classes int) *UnionFind {
	uf := new(UnionFind)
	uf.num_classes = num_classes
	uf.class = make([]int, num_classes)
	uf.class_size = make([]int, num_classes)

	for i := range num_classes {
		uf.class[i] = i
		uf.class_size[i] = 1
	}

	return uf
}

func (uf UnionFind) GetClass(n int) int {
	curr := n
	next := uf.class[curr]

	for curr != next {
		temp := uf.class[next]
		uf.class[curr] = temp
		curr = next
		next = temp
	}
	return curr
}

func (uf *UnionFind) Merge(a int, b int) {
	a_cl := uf.GetClass(a)
	b_cl := uf.GetClass(b)

	if a_cl != b_cl {
		uf.num_classes -= 1
		if uf.class_size[a_cl] < uf.class_size[b_cl] {
			uf.class[a_cl] = b_cl
			uf.class_size[b_cl] += uf.class_size[a_cl]
			uf.class_size[a_cl] = 0
		} else {
			uf.class[b_cl] = a_cl
			uf.class_size[a_cl] += uf.class_size[b_cl]
			uf.class_size[b_cl] = 0
		}
	}
}

func (uf *UnionFind) Finalize() {
	for i := range len(uf.class) {
		uf.GetClass(i)
	}
}
