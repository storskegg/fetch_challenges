package main

import (
	"sort"
	"strings"
)

type Ledger map[rune]int

func IsPyramidWord(word string) (b bool) {
	// Let's be case-insensitive
	word = strings.ToLower(word)

	ledger, lenLedger := LedgerFromWord(word)

	lastIdx := lenLedger - 1
	counts := SortLedgerCounts(ledger, lenLedger)

	return IsCardinalIncrementalOrder(counts, lastIdx)
}

func LedgerFromWord(word string) (*Ledger, int) {
	ledger := make(Ledger, len(word))

	for _, letter := range word {
		_, ok := ledger[letter]
		if ok {
			ledger[letter]++
		} else {
			ledger[letter] = 1
		}
	}

	return &ledger, len(ledger)
}

func SortLedgerCounts(ledger *Ledger, lenLedger int) []int {
	vals := make([]int, lenLedger)
	c := 0
	for _, n := range *ledger {
		vals[c] = n
		c++
	}
	sort.Ints(vals)
	return vals
}

func IsCardinalIncrementalOrder(counts []int, lastIdx int) (b bool) {
	for idx, n := range counts {
		if idx == 0 {
			if n == 1 {
				continue
			}

			break
		}

		if n == (counts[idx-1] + 1) {
			if idx == lastIdx {
				b = true
			}
			continue
		}
		break
	}
	return
}
