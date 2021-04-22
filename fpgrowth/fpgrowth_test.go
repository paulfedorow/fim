package fpgrowth

import (
	"fim/ints"
	"sort"
	"testing"
)

func TestMine(t *testing.T) {
	var txs = []Itemset {
		{1, 3},
		{1, 2, 5},
		{1, 2, 3},
		{4, 3, 2, 1},
		{3, 4},
	}
	var expectedFreqItemsets = []Itemset {
		{1}, {2}, {3}, {4},
		{1, 2}, {1, 3}, {2, 3}, {3, 4},
		{1, 2, 3},
	}
	var freqItemsets = Mine(txs, 2)
	for _, itemset := range freqItemsets {
		sort.Ints(itemset)
	}
	sort.Slice(freqItemsets, func(i, j int) bool { return ints.Less(freqItemsets[i], freqItemsets[j]) })
	if len(freqItemsets) == len(expectedFreqItemsets) {
		for i, expected := range expectedFreqItemsets {
			if !ints.Equals(freqItemsets[i], expected) {
				t.Fail()
			}
		}
	} else {
		t.Fail()
	}
}
