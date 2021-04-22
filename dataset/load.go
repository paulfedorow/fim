package dataset

import (
	"bufio"
	"fim/mine"
	"os"
	"strconv"
	"strings"
)

// Load reads the transactions that are contained in the given file.
//
// The expected format of the given file is as follows:
//  File = {Transaction}
//  Transaction = Item {" " Item} "\n"
//  Item = {"0" ... "9"}
//
// An example file could look like this:
//  1 2 3
//  3
//  1 2
func Load(input *os.File) []mine.Itemset {
	var scanner = bufio.NewScanner(input)
	var txs []mine.Itemset
	for scanner.Scan() {
		var itemTexts = strings.Split(strings.Trim(scanner.Text(), " \r\n"), " ")
		var tx []int
		for _, itemText := range itemTexts {
			if item, err := strconv.ParseInt(itemText, 10, 64); err == nil {
				tx = append(tx, int(item))
			}
		}
		txs = append(txs, tx)
	}
	return txs
}

