package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	defVerRelationship = "equal to"
)

func VersionStringCompare(versionA, versionB string) (string, error) {
	fmt.Printf("Comparing %s and %s...\n", versionA, versionB)
	normalizedA := versionA
	normalizedB := versionB

	pA := strings.Count(versionA, ".")
	pB := strings.Count(versionB, ".")

	fmt.Printf("pA: %d\n", pA)
	fmt.Printf("pB: %d\n", pB)

	if pA > pB {
		normalizedA = versionA
		normalizedB = versionB + strings.Repeat(".0", pA-pB)
	} else if pB > pA {
		normalizedA = versionA + strings.Repeat(".0", pB-pA)
		normalizedB = versionB
	}

	fmt.Println("normalizedA: " + normalizedA)
	fmt.Println("normalizedB: " + normalizedB)

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

	return fmt.Sprintf("%s is %s %s\n", versionA, relationship, versionB), nil
}
