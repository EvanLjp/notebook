package main

import "github.com/evanljp/algorithm/tags"

var _ = []tags.LUT{tags.LookupTable}

func k(a, b []int) float64 {
	return float64(a[1]-b[1]) / float64(a[0]-b[0])
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	max := 0
	for i, p1 := range points {
		dict := make(map[float64]int)
		for j, p2 := range points {
			if i != j {
				dict[k(p1, p2)]++
			}
		}
		for _, cnt := range dict {
			if cnt+1 > max {
				max = cnt + 1
			}
		}
	}
	return max

}

func main() {

}
