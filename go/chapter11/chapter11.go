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

type BloomFilter struct {
	num_entries uint64
	table       []uint64
	hash_mult   []uint64
	hash_add    []uint64
}

func BloomFilterInit(num_entries uint64) *BloomFilter {
	x := new(BloomFilter)
	x.num_entries = num_entries

	table_num_words := (num_entries + 63) / 64
	x.table = make([]uint64, table_num_words)

	num_hash := 3
	x.hash_mult = make([]uint64, num_hash)
	x.hash_add = make([]uint64, num_hash)

	for i := range num_hash {
		x.hash_add[i] = uint64(i + 1)
		x.hash_mult[i] = uint64(2*i + 3)
	}

	return x
}

func (filter *BloomFilter) Add(v uint64) {
	num_hash := len(filter.hash_add)
	for i := range num_hash {
		e := (filter.hash_mult[i] * (filter.hash_add[i] + v)) % filter.num_entries

		word := e / 64
		bit := e & 63

		// fmt.Println("AAA", word, bit)

		filter.table[word] |= uint64(1) << bit
	}
}

func (filter BloomFilter) Check(v uint64) bool {
	num_hash := len(filter.hash_add)
	for i := range num_hash {
		e := (filter.hash_mult[i] * (filter.hash_add[i] + v)) % filter.num_entries

		word := e / 64
		bit := e & 63

		// fmt.Println("BBB", word, bit)

		q := (filter.table[word] >> bit) & uint64(1)
		if q == 0 {
			return false
		}
	}

	return true
}
