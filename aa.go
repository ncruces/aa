// Package aa implements immutable AA trees.
package aa

import "cmp"

// Tree is an immutable AA tree,
// a form of self-balancing binary search tree.
//
// Use *Tree as a reference type; the zero value for *Tree (nil) is the empty tree:
//
//	var empty *aa.Tree[int, string]
//	one := empty.Put(1, "one")
//	one.Has(1) ⟹ true
//
// Note: the zero value for Tree{} is a valid — but non-empty — tree.
type Tree[K cmp.Ordered, V any] struct {
	left  *Tree[K, V]
	right *Tree[K, V]
	key   K
	value V
	balance
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
	return tree.balance.level()
}

// Len returns the number of nodes in this tree.
func (tree *Tree[K, V]) Len() int {
	if tree == nil {
		return 0
	}
	return tree.balance.len()
}

// Min finds the least key in this tree,
// and returns the node for that key,
// or nil if this tree is empty.
func (tree *Tree[K, V]) Min() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	for tree.left != nil {
		tree = tree.left
	}
	return tree
}

// Max finds the greatest key in this tree,
// and returns the node for that key,
// or nil if this tree is empty.
func (tree *Tree[K, V]) Max() *Tree[K, V] {
	if tree == nil {
		return nil
	}
	for tree.right != nil {
		tree = tree.right
	}
	return tree
}

// Floor finds the greatest key in this tree less-than or equal-to key,
// and returns the node for that key,
// or nil if no such key exists in this tree.
func (tree *Tree[K, V]) Floor(key K) *Tree[K, V] {
	var node *Tree[K, V]
	for tree != nil {
		if cmp.Less(key, tree.key) {
			tree = tree.left
		} else {
			node = tree
			tree = tree.right
		}
	}
	return node
}

// Ceil finds the least key in this tree greater-than or equal-to key,
// and returns the node for that key,
// or nil if no such key exists in this tree.
func (tree *Tree[K, V]) Ceil(key K) *Tree[K, V] {
	var node *Tree[K, V]
	for tree != nil {
		if cmp.Less(tree.key, key) {
			tree = tree.right
		} else {
			node = tree
			tree = tree.left
		}
	}
	return node
}

// Get retrieves the value for a given key;
// found indicates whether key exists in this tree.
func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	// Floor uses 2-way search, which is faster for strings:
	//   https://go.dev/issue/71270
	//   https://user.it.uu.se/~arnea/ps/searchproc.pdf
	// Both Floor/Ceil work; Floor is faster since AA trees lean right.
	node := tree.Floor(key)
	if node != nil && (key == node.key || key != key) {
		return node.value, true
	}
	return // zero, false
}

// Has reports whether key exists in this tree.
func (tree *Tree[K, V]) Has(key K) bool {
	_, found := tree.Get(key)
	return found
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
//	tree.Add(key).Has(key) ⟹ true
func (tree *Tree[K, V]) Add(key K) *Tree[K, V] {
	return tree.Patch(key, func(node *Tree[K, V]) (value V, ok bool) {
		return value, node == nil
	})
}

// Patch finds key in this tree, calls update with the node for that key
// (or nil, if key is not found), and returns a (possibly) modified tree.
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

	switch cmp.Compare(key, tree.key) {
	default:
		if value, ok := update(tree); ok {
			copy := *tree
			copy.value = value
			return &copy
		}
		return tree

	case -1:
		left := tree.left.Patch(key, update)
		if left == tree.left {
			return tree
		}
		copy := *tree
		copy.left = left
		return copy.ins_rebalance()

	case +1:
		right := tree.right.Patch(key, update)
		if right == tree.right {
			return tree
		}
		copy := *tree
		copy.right = right
		return copy.ins_rebalance()
	}
}

// Delete returns a (possibly) modified tree with key removed from it.
// The optional pred is called to confirm deletion.
//
//	tree.Delete(key).Has(key) ⟹ false
func (tree *Tree[K, V]) Delete(key K, pred ...func(node *Tree[K, V]) bool) *Tree[K, V] {
	var p func(*Tree[K, V]) bool
	if len(pred) > 0 {
		p = pred[0]
	}
	return tree.delete(key, p)
}

func (tree *Tree[K, V]) delete(key K, pred func(node *Tree[K, V]) bool) *Tree[K, V] {
	if tree == nil {
		return nil
	}

	switch cmp.Compare(key, tree.key) {
	case -1:
		left := tree.left.delete(key, pred)
		if left == tree.left {
			return tree
		}
		copy := *tree
		copy.left = left
		return copy.del_rebalance()

	case +1:
		right := tree.right.delete(key, pred)
		if right == tree.right {
			return tree
		}
		copy := *tree
		copy.right = right
		return copy.del_rebalance()

	default:
		if pred != nil && !pred(tree) {
			return tree
		}

		// If tree.right is nil, tree.left is too.
		if tree.left == nil {
			return tree.right
		}

		copy := *tree
		var heir *Tree[K, V]
		// Either works; this saves a few allocs.
		if copy.Level() == copy.right.Level() {
			copy.right, heir = copy.right.DeleteMin()
		} else {
			copy.left, heir = copy.left.DeleteMax()
		}
		copy.key = heir.key
		copy.value = heir.value
		return copy.del_rebalance()
	}
}

// DeleteMin returns a modified tree with its least key removed from it,
// and the removed node.
func (tree *Tree[K, V]) DeleteMin() (_, node *Tree[K, V]) {
	if tree == nil {
		return nil, nil
	}
	if tree.left == nil {
		return tree.right, tree
	}
	copy := *tree
	copy.left, node = tree.left.DeleteMin()
	return copy.del_rebalance(), node
}

// DeleteMax returns a modified tree with its greatest key removed from it,
// and the removed node.
func (tree *Tree[K, V]) DeleteMax() (_, node *Tree[K, V]) {
	if tree == nil {
		return nil, nil
	}
	if tree.right == nil {
		return nil, tree // tree.left, tree
	}
	copy := *tree
	copy.right, node = tree.right.DeleteMax()
	return copy.del_rebalance(), node
}
