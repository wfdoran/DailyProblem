package chapter7

type TrieNode struct {
	r     rune
	end   bool
	root  bool
	value int
	m     map[rune](*TrieNode)
}

func (node *TrieNode) Insert(s string) {
	node.InsertValue(s, 0)
}

func (node *TrieNode) InsertValue(s string, value int) {
	curr := node

	for _, r := range s {
		next, ok := curr.m[r]

		if ok {
			curr = next
		} else {
			next := new(TrieNode)
			next.r = r
			next.root = false
			next.end = false
			next.value = 0
			next.m = make(map[rune]*TrieNode)

			curr.m[r] = next
			curr = next
		}
	}
	curr.end = true
	curr.value = value
}

func (node TrieNode) _walk(prefix string, curr string) string {
	s := prefix

	if !node.root {
		s += string(node.r)
		if node.end {
			curr += s + "\n"
		}
	}

	for _, sub := range node.m {
		curr = sub._walk(s, curr)
	}

	return curr
}

func (node *TrieNode) _walk2(s string) []string {
	var rv []string

	if node.end {
		rv = append(rv, s)
	}

	for ch, sub_node := range node.m {
		sub_s := s + string(ch)
		q := sub_node._walk2(sub_s)

		rv = append(rv, q...)
	}
	return rv
}

func (node *TrieNode) _walk3() int {
	if node == nil {
		return 0
	}

	rv := node.value

	for _, sub_node := range node.m {
		rv += sub_node._walk3()
	}
	return rv
}

func (node *TrieNode) String() string {
	return node._walk("", "")
}

func TrieRoot() *TrieNode {
	n := new(TrieNode)
	n.m = make(map[rune]*TrieNode)
	n.root = true
	n.value = 0
	return n
}

func (node TrieNode) Find(s string) bool {
	curr := node

	for _, r := range s {
		next, ok := curr.m[r]
		if !ok {
			return false
		}
		curr = *next
	}

	return curr.end
}

func (node *TrieNode) AutoComplete(s string) []string {
	curr := node

	for _, c := range s {
		next, ok := curr.m[c]
		if !ok {
			return nil
		}
		curr = next
	}

	return curr._walk2(s)
}

func (node *TrieNode) Sum(s string) int {
	curr := node

	for _, c := range s {
		next, ok := curr.m[c]
		if !ok {
			return 0
		}
		curr = next
	}

	return curr._walk3()
}
