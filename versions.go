package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	defVerRelationship = "equal to"
)

// VersionStringCompare is a method that accepts two version strings, and returns a relationship phrase and an error.
func VersionStringCompare(versionA, versionB string) (string, error) {
	fmt.Printf("Comparing %s and %s...\n", versionA, versionB)

	normalizedA, normalizedB := NormalizeVersionStrings(versionA, versionB)

	vA := strings.Split(normalizedA, ".")
	vB := strings.Split(normalizedB, ".")

	relationship := defVerRelationship

	for idx, valA := range vA {
		nA, err := strconv.Atoi(valA)
		if err != nil {
			return "", err
		}
		nB, err := strconv.Atoi(vB[idx])
		if err != nil {
			return "", err
		}

		if nA == nB {
			continue
		}

		if nA < nB {
			relationship = "before"
			break
		}

		if nA > nB {
			relationship = "after"
			break
		}
	}

	return fmt.Sprintf("%s is %s %s", versionA, relationship, versionB), nil
}

// NormalizeVersionStrings normalizes version strings so that `2` can be compared with `2.0.0.0.0` more easily
func NormalizeVersionStrings(versionA, versionB string) (normalizedA, normalizedB string) {
	normalizedA = versionA
	normalizedB = versionB

	pA := strings.Count(versionA, ".")
	pB := strings.Count(versionB, ".")

	if pA > pB {
		normalizedA = versionA
		normalizedB = versionB + strings.Repeat(".0", pA-pB)
	} else if pB > pA {
		normalizedA = versionA + strings.Repeat(".0", pB-pA)
		normalizedB = versionB
	}

	return
}
