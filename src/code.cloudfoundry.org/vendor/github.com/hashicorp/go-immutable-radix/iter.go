package iradix

import (
	"bytes"
)

// Iterator is used to iterate over a set of nodes
// in pre-order
type Iterator struct {
	node  *Node
	stack []edges
}

// SeekPrefixWatch is used to seek the iterator to a given prefix
// and returns the watch channel of the finest granularity
func (i *Iterator) SeekPrefixWatch(prefix []byte) (watch <-chan struct{}) {
	// Wipe the stack
	i.stack = nil
	n := i.node
	watch = n.mutateCh
	search := prefix
	for {
<<<<<<< HEAD
		// Check for key exhaution
=======
		// Check for key exhaustion
>>>>>>> Add locking + restart-on-failed-lock functionality
		if len(search) == 0 {
			i.node = n
			return
		}

		// Look for an edge
		_, n = n.getEdge(search[0])
		if n == nil {
			i.node = nil
			return
		}

		// Update to the finest granularity as the search makes progress
		watch = n.mutateCh

		// Consume the search prefix
		if bytes.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]

		} else if bytes.HasPrefix(n.prefix, search) {
			i.node = n
			return
		} else {
			i.node = nil
			return
		}
	}
}

// SeekPrefix is used to seek the iterator to a given prefix
func (i *Iterator) SeekPrefix(prefix []byte) {
	i.SeekPrefixWatch(prefix)
}

func (i *Iterator) recurseMin(n *Node) *Node {
	// Traverse to the minimum child
	if n.leaf != nil {
		return n
	}
<<<<<<< HEAD
	if len(n.edges) > 0 {
		// Add all the other edges to the stack (the min node will be added as
		// we recurse)
		i.stack = append(i.stack, n.edges[1:])
=======
	nEdges := len(n.edges)
	if nEdges > 1 {
		// Add all the other edges to the stack (the min node will be added as
		// we recurse)
		i.stack = append(i.stack, n.edges[1:])
	}
	if nEdges > 0 {
>>>>>>> Add locking + restart-on-failed-lock functionality
		return i.recurseMin(n.edges[0].node)
	}
	// Shouldn't be possible
	return nil
}

// SeekLowerBound is used to seek the iterator to the smallest key that is
// greater or equal to the given key. There is no watch variant as it's hard to
// predict based on the radix structure which node(s) changes might affect the
// result.
func (i *Iterator) SeekLowerBound(key []byte) {
	// Wipe the stack. Unlike Prefix iteration, we need to build the stack as we
	// go because we need only a subset of edges of many nodes in the path to the
<<<<<<< HEAD
	// leaf with the lower bound.
	i.stack = []edges{}
	n := i.node
	search := key

	found := func(n *Node) {
		i.node = n
		i.stack = append(i.stack, edges{edge{node: n}})
	}

=======
	// leaf with the lower bound. Note that the iterator will still recurse into
	// children that we don't traverse on the way to the reverse lower bound as it
	// walks the stack.
	i.stack = []edges{}
	// i.node starts off in the common case as pointing to the root node of the
	// tree. By the time we return we have either found a lower bound and setup
	// the stack to traverse all larger keys, or we have not and the stack and
	// node should both be nil to prevent the iterator from assuming it is just
	// iterating the whole tree from the root node. Either way this needs to end
	// up as nil so just set it here.
	n := i.node
	i.node = nil
	search := key

	found := func(n *Node) {
		i.stack = append(i.stack, edges{edge{node: n}})
	}

	findMin := func(n *Node) {
		n = i.recurseMin(n)
		if n != nil {
			found(n)
			return
		}
	}

>>>>>>> Add locking + restart-on-failed-lock functionality
	for {
		// Compare current prefix with the search key's same-length prefix.
		var prefixCmp int
		if len(n.prefix) < len(search) {
			prefixCmp = bytes.Compare(n.prefix, search[0:len(n.prefix)])
		} else {
			prefixCmp = bytes.Compare(n.prefix, search)
		}

		if prefixCmp > 0 {
			// Prefix is larger, that means the lower bound is greater than the search
			// and from now on we need to follow the minimum path to the smallest
			// leaf under this subtree.
<<<<<<< HEAD
			n = i.recurseMin(n)
			if n != nil {
				found(n)
			}
=======
			findMin(n)
>>>>>>> Add locking + restart-on-failed-lock functionality
			return
		}

		if prefixCmp < 0 {
			// Prefix is smaller than search prefix, that means there is no lower
			// bound
			i.node = nil
			return
		}

		// Prefix is equal, we are still heading for an exact match. If this is a
<<<<<<< HEAD
		// leaf we're done.
		if n.leaf != nil {
			if bytes.Compare(n.leaf.key, key) < 0 {
				i.node = nil
				return
			}
=======
		// leaf and an exact match we're done.
		if n.leaf != nil && bytes.Equal(n.leaf.key, key) {
>>>>>>> Add locking + restart-on-failed-lock functionality
			found(n)
			return
		}

<<<<<<< HEAD
		// Consume the search prefix
		if len(n.prefix) > len(search) {
			search = []byte{}
		} else {
			search = search[len(n.prefix):]
=======
		// Consume the search prefix if the current node has one. Note that this is
		// safe because if n.prefix is longer than the search slice prefixCmp would
		// have been > 0 above and the method would have already returned.
		search = search[len(n.prefix):]

		if len(search) == 0 {
			// We've exhausted the search key, but the current node is not an exact
			// match or not a leaf. That means that the leaf value if it exists, and
			// all child nodes must be strictly greater, the smallest key in this
			// subtree must be the lower bound.
			findMin(n)
			return
>>>>>>> Add locking + restart-on-failed-lock functionality
		}

		// Otherwise, take the lower bound next edge.
		idx, lbNode := n.getLowerBoundEdge(search[0])
		if lbNode == nil {
<<<<<<< HEAD
			i.node = nil
=======
>>>>>>> Add locking + restart-on-failed-lock functionality
			return
		}

		// Create stack edges for the all strictly higher edges in this node.
		if idx+1 < len(n.edges) {
			i.stack = append(i.stack, n.edges[idx+1:])
		}

<<<<<<< HEAD
		i.node = lbNode
=======
>>>>>>> Add locking + restart-on-failed-lock functionality
		// Recurse
		n = lbNode
	}
}

// Next returns the next node in order
func (i *Iterator) Next() ([]byte, interface{}, bool) {
	// Initialize our stack if needed
	if i.stack == nil && i.node != nil {
		i.stack = []edges{
			{
				edge{node: i.node},
			},
		}
	}

	for len(i.stack) > 0 {
		// Inspect the last element of the stack
		n := len(i.stack)
		last := i.stack[n-1]
		elem := last[0].node

		// Update the stack
		if len(last) > 1 {
			i.stack[n-1] = last[1:]
		} else {
			i.stack = i.stack[:n-1]
		}

		// Push the edges onto the frontier
		if len(elem.edges) > 0 {
			i.stack = append(i.stack, elem.edges)
		}

		// Return the leaf values if any
		if elem.leaf != nil {
			return elem.leaf.key, elem.leaf.val, true
		}
	}
	return nil, nil, false
}