package main

import (
	"sort"
	"strings"
)

// Ledger is a key/val map where the key is type rune, and the value is type int. it represents individual letters, and
// the count of their occurrences in a word.
type Ledger map[rune]int

// IsPyramidWord is a method that accepts a word, and returns whether it is a pyramid word.
func IsPyramidWord(word string) (b bool) {
	ledger := LedgerFromWord(word)
	counts := SortLedgerCounts(ledger)

	return IsCardinalIncrementalOrder(counts)
}

// LedgerFromWord is a method that accepts a word, and returns a Ledger containing each letter appearing in that word,
// and the number of their respective occurrences.
func LedgerFromWord(word string) Ledger {
	word = strings.ToLower(word) // let's be case-insensitive
	ledger := make(Ledger, len(word)) // minimize memory allocations

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

// SortLedgerCounts is a method that takes a Ledger's values, and returns them in ascending order
func SortLedgerCounts(ledger Ledger) []int {
	// minimize memory allocations, and allows us to write to the index rather than append the slice
	counts := make([]int, len(ledger))

	// range on a map gives us key/val pairs, so we need to keep track of an index for `counts` by hand
	c := 0
	for _, n := range ledger {
		counts[c] = n
		c++
	}
	sort.Ints(counts) // ascending order
	return counts
}

// IsCardinalIncrementalOrder is a method that determines if an int slice is cardinally in incremental order -- that is
// that it starts with 1, and each index increments the previous value by 1.
// for example:
//   [1, 2, 3] passes
//   [1, 2, 5] fails
//   [2, 3, 4] fails
//   [3, 2, 1] fails
func IsCardinalIncrementalOrder(counts []int) (b bool) {
	// if counts has a length of 1, and that value is 1, we need not go farther
	if len(counts) == 1 && counts[0] == 1 {
		return true
	}

	lastIdx := len(counts) - 1 // cache this for efficiency

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
