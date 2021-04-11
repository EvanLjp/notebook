package main

import (
	"github.com/evanljp/algorithm/tags"
	"sort"
)

var _ = []tags.LUT{tags.LookupTable}

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		dict[string(runes)] = append(dict[string(runes)], str)
	}

	var res [][]string
	for _, strings := range dict {
		res = append(res, strings)

	}
	return res

}

func main() {

}
