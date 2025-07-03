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

func (node *TrieNode) String() string {
	return node._walk("", "")
}

func TrieRoot() *TrieNode {
	n := new(TrieNode)
	n.m = make(map[rune]*TrieNode)
	n.root = true
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
