package apriori

import (
	"sort"
	"testing"
)

func ItemsetLess(itemset1 Itemset, itemset2 Itemset) bool {
	if len(itemset1) == len(itemset2) {
		for i := 0; i < len(itemset1); i += 1 {
			if itemset1[i] < itemset2[i] {
				return true
			}
		}
		return false
	} else {
		return len(itemset1) < len(itemset2)
	}
}

func ItemsetEqual(itemset1 Itemset, itemset2 Itemset) bool {
	return !ItemsetLess(itemset1, itemset2) && !ItemsetLess(itemset1, itemset2)
}

func TestApriori(t *testing.T) {
	var txs = []Itemset {
		{1, 3},
		{1, 2, 5},
		{1, 2, 3},
		{1, 2, 3, 4},
		{2, 4},
	}
	var expectedFreqItemsets = []Itemset {
		{1}, {2}, {3}, {4},
		{1, 2}, {1, 3}, {2, 3}, {3, 4},
		{1, 2, 3},
	}
	var freqItemsets = Mine(txs, 2)
	sort.Slice(freqItemsets, func (i, j int) bool {
		return ItemsetLess(freqItemsets[i], freqItemsets[j])
	})
	if len(freqItemsets) == len(expectedFreqItemsets) {
		for i, expected := range expectedFreqItemsets {
			if !ItemsetEqual(freqItemsets[i], expected) {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}
