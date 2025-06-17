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

	if key < self.key {
		self.left = self.left.Insert(key, data)
	} else {
		self.right = self.right.Insert(key, data)
	}

	self.UpdateHeight()

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
