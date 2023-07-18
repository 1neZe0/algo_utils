package main

import (
	"fmt"
	"github.com/hbollon/go-edlib"
)

func main() {
	// Fuzzy search
	strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting", "helo"}
	res, err := edlib.FuzzySearchThreshold("testnig", strList, 0.7, edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result for 'testnig': %s", res)
	}

	res, err = edlib.FuzzySearchThreshold("hello", strList, 0.7, edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result for 'hello': %s", res)
	}

	// Longest common sequence
	lcs := edlib.LCS("ABCD", "ACBAD")
	fmt.Printf("Length of their LCS: %d", lcs)

	res2, err := edlib.LCSBacktrack("ABCD", "ACBAD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("LCS: %s", res2)

	// Diff between two strings
	res3, err := edlib.LCSDiff("computer", "houseboat")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("LCS: \n%s\n%s", res3[0], res3[1])

	// Hamming distance of two equal string
	ham, err := edlib.HammingDistance("computer", "comsebof")
	fmt.Println(ham)
}
