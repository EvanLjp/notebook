package main

import (
	"fmt"
	"github.com/evanljp/algorithm/tags"
	"sort"
)

var _ = []tags.LUT{tags.LookupTable}

type charCount struct {
	char  rune
	count int
}

func frequencySort(s string) string {
	var charCounts [128]charCount
	for _, c := range s {
		charCounts[c].char = c
		charCounts[c].count++
	}
	sort.Slice(charCounts[:], func(i, j int) bool {
		return charCounts[i].count > charCounts[j].count
	})
	result := make([]rune, 0, len(s))
	for _, cc := range charCounts {
		// if cc.count == 0 { // optimisation
		// 	break
		// }
		for ; cc.count > 0; cc.count-- {
			result = append(result, cc.char)
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("%s", frequencySort("tree"))

}
