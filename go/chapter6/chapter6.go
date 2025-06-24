package chapter6

import (
	"cmp"
	"fmt"
)

type TreeNode[K cmp.Ordered, T any] struct {
	key    K
	data   T
	height int
	left   *TreeNode[K, T]
	right  *TreeNode[K, T]
}

func NewTreeNode[K cmp.Ordered, T any](key K, data T) *TreeNode[K, T] {
	n := new(TreeNode[K, T])
	n.key = key
	n.data = data
	n.height = 1
	n.left = nil
	n.right = nil

	return n
}

func (self *TreeNode[K, T]) UpdateHeight() {
	left_height := 0
	if self.left != nil {
		left_height = self.left.height
	}
	right_height := 0
	if self.right != nil {
		right_height = self.right.height
	}

	if left_height > right_height {
		self.height = left_height + 1
	} else {
		self.height = right_height + 1
	}
}

func (self *TreeNode[K, T]) Balance() *TreeNode[K, T] {
	left_height := 0
	if self.left != nil {
		left_height = self.left.height
	}
	right_height := 0
	if self.right != nil {
		right_height = self.right.height
	}
	if left_height > right_height {
		self.height = left_height + 1
	} else {
		self.height = right_height + 1
	}

	/*             self                        x
	              /    \                      /  \
	             x       C                   A   self
				/ \                              /   \
			   A	B		                    B     C
	*/

	if left_height > right_height+1 {
		x := self.left
		a := x.left
		b := x.right
		c := self.right

		x.left = a
		x.right = self
		self.left = b
		self.right = c

		self.UpdateHeight()
		x.UpdateHeight()
		return x
	}

	/*                 self                       x
		              /    \                     /  \
					A       x                  self  C
					       / \                 /  \
	                      B   C               A    B

	*/

	if left_height < right_height-1 {
		x := self.right
		a := self.left
		b := x.left
		c := x.right

		x.left = self
		x.right = c
		self.left = a
		self.right = b

		self.UpdateHeight()
		x.UpdateHeight()
		return x
	}

	return self
}

func (self *TreeNode[K, T]) Insert(key K, data T) *TreeNode[K, T] {
	if self == nil {
		return NewTreeNode(key, data)
	}

	if key == self.key {
		self.data = data
	} else if key < self.key {
		self.left = self.left.Insert(key, data)
	} else {
		self.right = self.right.Insert(key, data)
	}
	return self.Balance()
}

func (self *TreeNode[K, T]) Find(key K) (bool, T) {
	if self == nil {
		var bogus T
		return false, bogus
	}

	if self.key == key {
		return true, self.data
	}

	if key < self.key {
		return self.left.Find(key)
	} else {
		return self.right.Find(key)
	}
}

func (self *TreeNode[K, T]) String() string {
	rv := "( "

	if self.left != nil {
		rv += self.left.String()
		rv += " "
	}
	rv += fmt.Sprint(self.key)

	if self.right != nil {
		rv += " "
		rv += self.right.String()
	}
	rv += " )"

	return rv
}

func (self *TreeNode[K, T]) Remove(key K) (bool, T, *TreeNode[K, T]) {
	if self == nil {
		var bogus T
		return false, bogus, self
	}

	if key < self.key {
		ok, data, update := self.left.Remove(key)
		self.left = update
		ret := self.Balance()
		return ok, data, ret
	}

	if key > self.key {
		ok, data, update := self.right.Remove(key)
		self.right = update
		ret := self.Balance()
		return ok, data, ret
	}

	if self.right == nil {
		return true, self.data, self.left
	}

	if self.left == nil {
		return true, self.data, self.right
	}

	if self.right.left == nil {
		self.right.left = self.left
		ret := self.right.Balance()
		return true, self.data, ret
	}

	if self.left.right == nil {
		self.left.right = self.right
		ret := self.right.Balance()
		return true, self.data, ret
	}

	succ := self.right
	for ; succ.left != nil; succ = succ.left {
	}

	temp_key := self.key
	self.key = succ.key
	succ.key = temp_key

	temp_data := self.data
	self.data = succ.data
	succ.data = temp_data

	ok, data, right_ret := self.right.Remove(key)
	self.right = right_ret

	ret := self.Balance()

	return ok, data, ret
}
