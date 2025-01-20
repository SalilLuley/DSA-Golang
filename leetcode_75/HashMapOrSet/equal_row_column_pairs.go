package leetcode_75

import (
	"fmt"
	"strings"
)

func EqualRowColumnPairs() {
	// grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Printf("equalPairs(grid): %v\n", equalPairs(grid))
}
func equalPairs(grid [][]int) int {
	pairs := 0
	hm := make(map[string]int)
	for i := range grid {
		rowStr := sliceToString(grid[i])
		hm[rowStr]++
	}
	// Convert each column to a string and compare with the map
	for j := 0; j < len(grid); j++ {
		col := make([]int, len(grid))
		for i := range grid {
			col[i] = grid[i][j]
		}
		colStr := sliceToString(col)
		if v, ok := hm[colStr]; ok {
			pairs += v
		}
	}
	return pairs
}

func sliceToString(slice []int) string {
	strs := make([]string, len(slice))
	for i, v := range slice {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(strs, "")
}
