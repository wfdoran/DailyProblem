package chapter7

type TrieNode struct {
	r    rune
	end  bool
	root bool
	m    map[rune](*TrieNode)
}

func (node *TrieNode) Insert(s string) {
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
			next.m = make(map[rune]*TrieNode)

			curr.m[r] = next
			curr = next
		}
	}
	curr.end = true
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

func (node *TrieNode) String() string {
	return node._walk("", "")
}

func TrieRoot() *TrieNode {
	n := new(TrieNode)
	n.m = make(map[rune]*TrieNode)
	n.root = true
	return n
}
