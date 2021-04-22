package fpgrowth

import (
	"sort"
)

type Itemset []int

// Mine determines the frequent itemsets for the given transactions.
func Mine(txs []Itemset, minSupport int) []Itemset {
	// Count items.
	var itemCount = make(map[int]int)
	for _, tx := range txs {
		for _, item := range tx {
			itemCount[item] += 1
		}
	}

	// Determine frequent items.
	var freqItems = make(map[int]int)
	for item, count := range itemCount {
		if count >= minSupport {
			freqItems[item] = count
		}
	}

	// Create FP-tree out of the frequent items that are contained in transactions.
	var fpTree = fpTreeNew()
	for _, tx := range txs {
		var itemset Itemset
		for _, item := range tx {
			if _, ok := freqItems[item]; ok {
				itemset = append(itemset, item)
			}
		}
		sort.Slice(itemset, func(i, j int) bool {
			if freqItems[itemset[i]] == freqItems[itemset[j]] {
				return itemset[i] < itemset[j]
			} else {
				return freqItems[itemset[i]] > freqItems[itemset[j]]
			}
		})
		fpTree.insert(itemset, 1)
	}

	// Determine frequent itemsets by mining the FP-tree.
	var freqItemsets []Itemset
	fpTree.mine(minSupport, &freqItemsets)
	return freqItemsets
}

// fpTree is an implementation of the FP-tree data structure as described in the FP-growth paper by Han et al.
type fpTree struct {
	root     *fpNode
	itemHead map[int]*fpNode // indexed by item
}

// fpNode represents the node of an FP-tree.
type fpNode struct {
	item     int
	count    int
	parent   *fpNode
	children map[int]*fpNode // points to the child nodes indexed by their item
	itemNext *fpNode         // points to the next node that contains the same item
}

// fpTreeNew creates an empty FP-tree.
func fpTreeNew() *fpTree {
	return &fpTree{root: &fpNode{children: make(map[int]*fpNode)}, itemHead: make(map[int]*fpNode)}
}

// insert inserts the given itemset into the FP-tree with the given count.
func (t *fpTree) insert(itemset Itemset, count int) {
	var node = t.root
	for _, item := range itemset {
		if child, ok := node.children[item]; ok {
			// Child for the item already exists. Adjust the count and proceed.
			child.count += count
			node = child
		} else {
			// Child for the item does not exist yet. Create a new node, append it as child and proceed.
			var newNode = &fpNode{item: item, count: count, parent: node, children: make(map[int]*fpNode)}
			node.children[item] = newNode
			node = newNode
			if itemHead, ok := t.itemHead[item]; ok {
				newNode.itemNext = itemHead
			}
			t.itemHead[item] = newNode
		}
	}
}

// mine extracts the frequent itemsets out of the FP-tree.
func (t *fpTree) mine(minSupport int, freqItemsets *[]Itemset) {
	type stackEntry struct {
		fpTree *fpTree
		itemset Itemset
	}
	var stack = []stackEntry{{fpTree: t, itemset: nil}}
	for len(stack) > 0 {
		var fpTree = stack[0].fpTree
		var itemset = stack[0].itemset
		stack = stack[1:]
		for item, head := range fpTree.itemHead {
			// Determine the count of item in the fpTree.
			var count = 0
			var node = head
			for node != nil {
				count += node.count
				node = node.itemNext
			}
			if count >= minSupport {
				// Construct a frequent itemset by combining the item with the itemset from the stack.
				var newItemset = make(Itemset, len(itemset)+1)
				copy(newItemset, itemset)
				newItemset[len(newItemset)-1] = item
				*freqItemsets = append(*freqItemsets, newItemset)
				// Build an FP-tree conditioned by the item and add it to the stack.
				var conditionalTree = fpTree.conditionalTree(item, minSupport)
				stack = append(stack, stackEntry{fpTree: conditionalTree, itemset: newItemset})
			}
		}
	}
}

// conditionalTree creates an FP-tree conditioned by the given item.
func (t *fpTree) conditionalTree(item int, minSupport int) *fpTree {
	// Extract all prefixes of nodes that contain the conditional item.
	type prefix struct {
		itemset Itemset
		count int
	}
	var prefixes []prefix
	var itemCount = make(map[int]int)
	var node = t.itemHead[item]
	for node != nil {
		// Extract prefix by traversing upwards to the root node and collecting the encountered items.
		var count = node.count
		var parent = node.parent
		var itemset Itemset
		for parent != t.root {
			itemCount[parent.item] += count
			itemset = append(itemset, parent.item)
			parent = parent.parent
		}
		prefixes = append(prefixes, prefix{itemset: itemset, count: count})
		node = node.itemNext
	}

	// Build an FP-tree out of the prefixes.
	var fpTree = fpTreeNew()
	for _, prefix := range prefixes {
		var itemset Itemset
		for i, _ := range prefix.itemset {
			var item = prefix.itemset[len(prefix.itemset)-1-i] // the items in prefix were collected in reverse order
			if itemCount[item] >= minSupport {
				itemset = append(itemset, item)
			}
		}
		fpTree.insert(itemset, prefix.count)
	}
	return fpTree
}