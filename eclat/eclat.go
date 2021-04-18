package apriori

import (
	"fim/ints"
	"sort"
)

type Itemset []int

// Mine determines the frequent itemsets for the given transactions.
func Mine(txs []Itemset, minSupport int) []Itemset {
	// Create vertical representation of the transactions and count item pairs.
	var itemTidsets = make(map[int][]int)
	var itemPairCount = make(map[int]map[int]int)
	for tid, tx := range txs {
		for k, item1 := range tx {
			itemTidsets[item1] = append(itemTidsets[item1], tid)
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

	var freqItemsets []Itemset

	var atoms []atom

	// Determine frequent 1-itemsets and insert them into the atom set.
	for item, tidset := range itemTidsets {
		if len(tidset) >= minSupport {
			var itemset = Itemset{item}
			freqItemsets = append(freqItemsets, itemset)
			atoms = append(atoms, atom{itemset: itemset, tidset: tidset})
		}
	}

	// Sort the atoms by the order of increasing tidset size. This reduces the number of generated atoms.
	sort.Slice(atoms, func (i, j int) bool { return len(atoms[i].tidset) < len(atoms[j].tidset) })

	// The first level of the eclat function is inlined here, so that we can use the item pairs to accelerate the
	// support checking.
	for k, atom1 := range atoms {
		var newAtoms []atom
		for _, atom2 := range atoms[k+1:] {
			var item1, item2 = ints.MinMax(atom1.itemset[0], atom2.itemset[0])
			if counts, ok := itemPairCount[item1]; ok {
				if count, ok := counts[item2]; ok {
					if count >= minSupport {
						var itemset = Itemset{item1, item2}
						freqItemsets = append(freqItemsets, itemset)
						newAtoms = append(newAtoms, atom{
							itemset: itemset,
							tidset:  ints.Intersection(atom1.tidset, atom2.tidset),
						})
					}
				}
			}
		}
		eclat(newAtoms, minSupport, &freqItemsets)
	}

	return freqItemsets
}

type atom struct {
	itemset Itemset
	tidset []int
}

// eclat performs the eclat algorithm on the given atoms and collects any frequent itemsets into freqItemsets.
func eclat(atoms []atom, minSupport int, freqItemsets *[]Itemset) {
	// Perform the eclat algorithm by recursively combining atoms to larger itemsets.
	for k, atom1 := range atoms {
		var newAtoms []atom
		for _, atom2 := range atoms[k+1:] {
			var tidset = ints.Intersection(atom1.tidset, atom2.tidset)
			if len(tidset) >= minSupport {
				var itemset = ints.Union(atom1.itemset, atom2.itemset)
				*freqItemsets = append(*freqItemsets, itemset)
				newAtoms = append(newAtoms, atom{itemset: itemset, tidset: tidset})
			}
		}
		eclat(newAtoms, minSupport, freqItemsets)
	}
}
