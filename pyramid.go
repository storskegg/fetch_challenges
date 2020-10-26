package main

import (
	"sort"
	"strings"
)

type Ledger map[rune]int

func IsPyramidWord(word string) (b bool) {
	ledger := LedgerFromWord(word)
	counts := SortLedgerCounts(ledger)

	return IsCardinalIncrementalOrder(counts)
}

func LedgerFromWord(word string) Ledger {
	word = strings.ToLower(word)
	ledger := make(Ledger, len(word))

	for _, letter := range word {
		_, ok := ledger[letter]
		if ok {
			ledger[letter]++
		} else {
			ledger[letter] = 1
		}
	}

	return ledger
}

func SortLedgerCounts(ledger Ledger) []int {
	lenLedger := len(ledger)
	vals := make([]int, lenLedger)
	c := 0
	for _, n := range ledger {
		vals[c] = n
		c++
	}
	sort.Ints(vals)
	return vals
}

func IsCardinalIncrementalOrder(counts []int) (b bool) {
	if len(counts) == 1 && counts[0] == 1 {
		return true
	}

	lastIdx := len(counts) - 1

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
