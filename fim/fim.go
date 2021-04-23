package main

import (
	"fim/dataset"
	"fim/mine"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	// Parse arguments.
	var inputPath = flag.String("i", "datasets/retail.dat", "input path.")
	var outputPath = flag.String("o", "-", "output path.")
	var algorithm = flag.String("a", "apriori", "algorithm: apriori, eclat or fpgrowth.")
	var support = flag.Float64("s", 0.1, "minimal support.")
	flag.Parse()

	// Read dataset.
	input, err := os.Open(*inputPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	defer func() { _ = input.Close() }()
	var txs = dataset.Load(input)

	// Determine the frequent itemsets of the given dataset.
	var minSupport = int(math.Ceil(*support * float64(len(txs))))
	var freqItemsets []mine.Itemset
	switch *algorithm {
	case "apriori":
		freqItemsets = mine.Apriori(txs, minSupport)
	case "eclat":
		freqItemsets = mine.Eclat(txs, minSupport)
	case "fpgrowth":
		freqItemsets = mine.FPGrowth(txs, minSupport)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "unknown algorithm '%s'", *algorithm)
		os.Exit(1)
	}

	// Write the frequent itemsets into the output file.
	var output *os.File
	if *outputPath == "-" {
		output = os.Stdout
	} else {
		output, err = os.Open(*outputPath)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		defer func() { _ = input.Close() }()
	}
	for _, itemset := range freqItemsets {
		sort.Ints(itemset)
		for i, item := range itemset {
			if i < len(itemset)-1 {
				_, _ = fmt.Fprintf(output, "%d ", item)
			} else {
				_, _ = fmt.Fprintf(output, "%d\n", item)
			}
		}
	}
}
