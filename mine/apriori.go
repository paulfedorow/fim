package mine

import (
	"fim/ints"
	"sort"
)

// MineApriori determines the frequent itemsets for the given transactions.
func MineApriori(txs []Itemset, minSupport int) []Itemset {
	// Count items and item pairs.
	var itemCount = make(map[int]int)
	var itemPairCount = make(map[int]map[int]int)
	for _, tx := range txs {
		for k, item1 := range tx {
			itemCount[item1] += 1
			for _, item2 := range tx[k+1:] {
				item1, item2 = ints.MinMax(item1, item2)
				if _, ok := itemPairCount[item1]; !ok {
					itemPairCount[item1] = make(map[int]int)
				}
				itemPairCount[item1][item2] += 1
			}
		}
		sort.Ints(tx)
	}

	var allFreqItemsets []Itemset

	// Determine frequent 1-itemsets.
	for item, count := range itemCount {
		if count >= minSupport {
			allFreqItemsets = append(allFreqItemsets, Itemset{item})
		}
	}

	var prevFreqItemsets []Itemset

	// Determine frequent 2-itemsets.
	for item1, counts := range itemPairCount {
		for item2, count := range counts {
			if count >= minSupport {
				item1, item2 = ints.MinMax(item1, item2)
				allFreqItemsets = append(allFreqItemsets, Itemset{item1, item2})
				prevFreqItemsets = append(prevFreqItemsets, Itemset{item1, item2})
			}
		}
	}

	// Iteratively determine larger frequent itemsets by generating candidates out of the previously found frequent
	// itemsets.
	for len(prevFreqItemsets) > 0 {
		var candidates = candidates(prevFreqItemsets)
		prevFreqItemsets = nil
		if !candidates.empty() {
			for _, tx := range txs {
				candidates.count(tx)
			}
			prevFreqItemsets = candidates.mine(minSupport)
			allFreqItemsets = append(allFreqItemsets, prevFreqItemsets...)
		}
	}

	return allFreqItemsets
}

// candidates generates larger frequent itemsets by combining previously found frequent itemsets.
func candidates(prevFreqItemsets []Itemset) *trie {
	var trie = trieNew()
	for k, itemset1 := range prevFreqItemsets {
		for _, itemset2 := range prevFreqItemsets[k+1:] {
			var samePrefix = true
			for i, _ := range itemset1[:len(itemset1)-1] {
				if itemset1[i] != itemset2[i] {
					samePrefix = false
					break
				}
			}
			if samePrefix {
				// Generate a candidate by combining itemset1 and itemset2. The items of the candidate are sorted, this
				// results in a more compact trie.
				var candidate = make(Itemset, len(itemset1)+1)
				if itemset1[len(itemset1)-1] < itemset2[len(itemset2)-1] {
					copy(candidate, itemset1)
					candidate[len(candidate)-1] = itemset2[len(itemset2)-1]
				} else {
					copy(candidate, itemset2)
					candidate[len(candidate)-1] = itemset1[len(itemset1)-1]
				}
				// Insert the candidate into the candidate trie. The trie will be used to find the candidates that are
				// frequent.
				trie.insert(candidate)
			}
		}
	}
	return trie
}

// trie is an implementation of the trie data structure. It contains the candidate itemsets and keeps track of their
// frequency.
type trie struct {
	// Children nodes, indexed by item.
	children map[int]*trie

	// Frequency of the itemset that consists out of the items, that are on the path from the root to this node.
	frequency int
}

// trieNew creates an empty trie.
func trieNew() *trie {
	return &trie{children: make(map[int]*trie), frequency: 0}
}

// insert inserts the given itemset into the trie.
func (t *trie) insert(itemset Itemset) {
	var trie = t
	for _, item := range itemset {
		if _, ok := trie.children[item]; !ok {
			var newTrie = trieNew()
			trie.children[item] = newTrie
			trie = newTrie
		} else {
			trie = trie.children[item]
		}
	}
}

// empty returns true if the trie is empty, false otherwise.
func (t *trie) empty() bool {
	return len(t.children) == 0
}

// count increments the frequencies of the candidates contained in the trie if they are contained in the given
// transaction.
func (t *trie) count(tx Itemset) {
	// Create a set out of the items that are contained in the transaction.
	var txItems = make(map[int]struct{})
	for _, item := range tx {
		txItems[item] = struct{}{}
	}

	// Increment the frequencies of the nodes that are reachable from the root by following only the edges with items
	// that are contained in txItems.
	var stack = []*trie{t}
	for len(stack) > 0 {
		var trie = stack[0]
		stack = stack[1:]
		trie.frequency += 1
		for item, childTrie := range trie.children {
			if _, ok := txItems[item]; ok {
				stack = append(stack, childTrie)
			}
		}
	}
}

// mine determines the frequent itemsets that are contained in the trie.
func (t *trie) mine(minSupport int) []Itemset {
	type stackEntry struct {
		trie    *trie
		itemset Itemset
	}

	var freqItemsets []Itemset

	// Determine the frequent itemsets by finding all candidates that have frequencies which exceed minSupport.
	var stack = []stackEntry{{trie: t, itemset: nil}}
	for len(stack) > 0 {
		var trie = stack[0].trie
		var itemset = stack[0].itemset
		stack = stack[1:]
		if trie.empty() {
			// Leaf node reached. Collect the frequent itemset.
			freqItemsets = append(freqItemsets, itemset)
		} else {
			for item, childTrie := range trie.children {
				// Only follow edges that point to nodes with frequencies which exceed minSupport.
				if childTrie.frequency >= minSupport {
					var newItemset = make(Itemset, len(itemset)+1)
					copy(newItemset, itemset)
					newItemset[len(newItemset)-1] = item
					stack = append(stack, stackEntry{trie: childTrie, itemset: newItemset})
				}
			}
		}
	}

	return freqItemsets
}
