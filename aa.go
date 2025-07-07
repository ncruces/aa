// Package aa implements immutable AA trees.
package aa

import (
	"cmp"
	"iter"
)

// Tree is an immutable AA tree,
// a form of self-balancing binary search tree.
//
// Use *Tree as reference type; the zero value for *Tree (nil) is the empty tree:
//
//	var empty *aa.Tree[int, string]
//	one := empty.Put(1, "one")
//	one.Contains(1) ⟹ true
//
// Note: the zero value for Tree{} is a valid, but non-empty, tree.
type Tree[K cmp.Ordered, V any] struct {
	left  *Tree[K, V]
	right *Tree[K, V]
	key   K
	value V
	level int
}

// Key returns the key at the root of this tree.
//
// Note: getting the root key of an empty tree (nil)
// causes a runtime panic.
func (tree *Tree[K, V]) Key() K {
	return tree.key
}

// Value returns the value at the root of this tree.
//
// Note: getting the root value of an empty tree (nil)
// causes a runtime panic.
func (tree *Tree[K, V]) Value() V {
	return tree.value
}

// Left returns the left subtree of this tree,
// containing all keys less than its root key.
//
// Note: the left subtree of the empty tree is the empty tree (nil).
func (tree *Tree[K, V]) Left() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	return tree.left
}

// Right returns the right subtree of this tree,
// containing all keys greater than its root key.
//
// Note: the right subtree of the empty tree is the empty tree (nil).
func (tree *Tree[K, V]) Right() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	return tree.right
}

// Level returns the level of this AA tree.
//
// Notes:
//   - the level of the empty tree (nil) is 0.
//   - the height of a tree of level n is between n and 2·n.
//   - the number of nodes in a tree of level n is between 2ⁿ-1 and 3ⁿ-1.
func (tree *Tree[K, V]) Level() int {
	if tree == nil {
		return 0
	}
	return tree.level + 1
}

// Get retrieves the value for a given key;
// found indicates whether key exists in this tree.
func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	// For strings, 2-way search is faster:
	//   https://go.dev/issue/71270
	//   https://user.it.uu.se/~arnea/ps/searchproc.pdf
	var node *Tree[K, V]
	for tree != nil {
		if cmp.Less(key, tree.key) {
			tree = tree.left
		} else {
			node = tree
			tree = tree.right
		}
	}
	if node != nil && cmp.Compare(key, node.key) == 0 {
		return node.value, true
	}
	return // zero, false
}

// Contains reports whether key exists in this tree.
func (tree *Tree[K, V]) Contains(key K) bool {
	_, found := tree.Get(key)
	return found
}

// All returns an in-order iterator for this tree.
//
//	for node := range tree.All() {
//		fmt.Println(node.Key(), node.Value())
//	})
func (tree *Tree[K, V]) All() iter.Seq[*Tree[K, V]] {
	return func(yield func(*Tree[K, V]) bool) { tree.pull(yield) }
}

func (tree *Tree[K, V]) pull(yield func(*Tree[K, V]) bool) bool {
	if tree == nil {
		return true
	}
	return tree.left.pull(yield) && yield(tree) && tree.right.pull(yield)
}

// Put returns a modified tree with key set to value.
//
//	tree.Put(key, value).Get(key) ⟹ (value, true)
func (tree *Tree[K, V]) Put(key K, value V) *Tree[K, V] {
	return tree.Patch(key, func(*Tree[K, V]) (V, bool) {
		return value, true
	})
}

// Add returns a (possibly) modified tree that contains key.
//
//	tree.Add(key).Contains(key) ⟹ true
func (tree *Tree[K, V]) Add(key K) *Tree[K, V] {
	return tree.Patch(key, func(node *Tree[K, V]) (value V, ok bool) {
		return value, node == nil
	})
}

// Patch looks for key in this tree, calls update with the node for that key
// (or nil, if key is not found), and returns a possibly modified tree.
//
// The update callback can opt to set/update the value for the key,
// by returning (value, true), or not, by returning false.
func (tree *Tree[K, V]) Patch(key K, update func(node *Tree[K, V]) (value V, ok bool)) *Tree[K, V] {
	if tree == nil {
		if value, ok := update(tree); ok {
			return &Tree[K, V]{key: key, value: value}
		}
		return nil
	}

	copy := *tree
	switch cmp.Compare(key, tree.key) {
	default:
		if value, ok := update(tree); ok {
			copy.value = value
			return &copy
		}
		return tree
	case -1:
		copy.left = tree.left.Patch(key, update)
		if copy.left == tree.left {
			return tree
		}
	case +1:
		copy.right = tree.right.Patch(key, update)
		if copy.right == tree.right {
			return tree
		}
	}
	return copy.ins_rebalance()
}

// Delete returns a modified tree with key removed from it.
//
//	tree.Delete(key).Contains(key) ⟹ false
func (tree *Tree[K, V]) Delete(key K) *Tree[K, V] {
	if tree == nil {
		return nil
	}

	copy := *tree
	switch cmp.Compare(key, tree.key) {
	case -1:
		copy.left = tree.left.Delete(key)
		if copy.left == tree.left {
			return tree
		}
	case +1:
		copy.right = tree.right.Delete(key)
		if copy.right == tree.right {
			return tree
		}
	default:
		if tree.left == nil {
			return tree.right
		}
		heir := tree.left
		for heir.right != nil {
			heir = heir.right
		}
		copy.key = heir.key
		copy.value = heir.value
		copy.left = tree.left.Delete(heir.key)
	}
	return copy.del_rebalance()
}

func (tree Tree[K, V]) ins_rebalance() *Tree[K, V] {
	if tree.need_raise() { // avoid 2 rotations and allocs
		tree.level++
		return &tree
	}
	return tree.skew().split()
}

func (tree Tree[K, V]) del_rebalance() *Tree[K, V] {
	var want int
	if tree.left != nil && tree.right != nil {
		want = 1 + min(tree.left.level, tree.right.level)
	}
	if tree.level > want {
		tree.level = want
		if tree.right != nil && tree.right.level > want {
			copy := *tree.right
			copy.level = want
			tree.right = &copy
		}
		return tree.skew_r().split_r()
	}
	return &tree
}

func (tree Tree[K, V]) need_skew() bool {
	return tree.left != nil && tree.left.level == tree.level
}

func (tree Tree[K, V]) need_split() bool {
	return tree.right != nil && tree.right.right != nil && tree.right.right.level == tree.level
}

func (tree Tree[K, V]) need_raise() bool {
	return tree.left != nil && tree.right != nil &&
		tree.left.level == tree.level &&
		tree.right.level == tree.level
}

func (tree Tree[K, V]) skew() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = &tree
		return &copy
	}
	return &tree
}

func (tree Tree[K, V]) skew_r() *Tree[K, V] {
	if tree.need_skew() {
		copy := *tree.left
		tree.left = copy.right
		copy.right = tree.skew_r()
		return &copy
	}
	if tree.right != nil && tree.right.need_skew() {
		copy := *tree.right
		tree.right = copy.skew_r()
	}
	return &tree
}

func (tree Tree[K, V]) split() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.left = &tree
		copy.level++
		return &copy
	}
	return &tree
}

func (tree Tree[K, V]) split_r() *Tree[K, V] {
	if tree.need_split() {
		copy := *tree.right
		tree.right = copy.left
		copy.right = copy.right.split_r()
		copy.left = &tree
		copy.level++
		return &copy
	}
	return &tree
}
