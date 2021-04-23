package main

import (
	"fim/dataset"
	"fim/mine"
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"time"
)

type algorithm struct {
	name         string
	mine         func(txs []mine.Itemset, minSupport int) []mine.Itemset
	lastDuration time.Duration
}

func main() {
	// Parse arguments.
	var inputPath = flag.String("i", "datasets/retail.dat", "input path")
	var timeout = flag.Int("t", 30, "timeout in seconds")
	var startSupport = flag.Float64("s", 0.1, "start support")
	var endSupport = flag.Float64("e", 0.05, "end support")
	var supportStep = flag.Float64("d", 0.05, "support step")
	flag.Parse()
	var timeoutDuration = time.Duration(float64(*timeout) * math.Pow10(9))

	// Read dataset.
	input, err := os.Open(*inputPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	var txs = dataset.Load(input)

	// Measure the algorithm runtime on the given dataset with decreasing minimal supports.
	var algorithms = []algorithm{
		{name: "apriori", mine: mine.Apriori},
		{name: "eclat", mine: mine.Eclat},
		{name: "fpgrowth", mine: mine.FPGrowth},
	}
	var minSupport = *startSupport
	var stop = false
	fmt.Printf("minSupport,apriori,eclat,fpgrowth\n")
	for !stop && minSupport >= (*endSupport - math.Pow(0.1, 10)) {
		stop = true
		fmt.Printf("%f", minSupport)
		for i, algorithm := range algorithms {
			if algorithm.lastDuration > timeoutDuration {
				// skip algorithms whose last run took longer than timeoutDuration
				fmt.Printf(",")
				continue
			}
			runtime.GC() // force full GC to level the playing field
			var start = time.Now()
			_ = algorithm.mine(txs, int(math.Ceil(minSupport * float64(len(txs)))))
			var duration = time.Since(start)
			algorithms[i].lastDuration = duration
			stop = false
			fmt.Printf(",%f", duration.Seconds())
		}
		fmt.Printf("\n")
		minSupport = minSupport - *supportStep
	}
}
