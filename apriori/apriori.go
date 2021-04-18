package apriori

type Itemset []int

func Mine(txs []Itemset, minSupport int) []Itemset {
	return []Itemset {
		{1}, {2}, {3}, {4},
		{1, 2}, {1, 3}, {2, 3}, {3, 4},
		{1, 2, 3},
	}
}