package main

import "fmt"

func solution(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	val := 1
	row := 0
	col := -1
	idx_inc := 1
	for loop_count := n; loop_count > 0; loop_count-- {
		for i := 0; i < loop_count; i++ {
			col += idx_inc
			ret[row][col] = val
			val++
		}
		for i := 0; i < loop_count-1; i++ {
			row += idx_inc
			ret[row][col] = val
			val++
		}
		idx_inc *= -1
	}

	return ret
}

// Compare2DSlices compares two 2-dimensional slices and returns true if they are equal.
func Compare2DSlices(slice1, slice2 [][]int) bool {
	// First, check if the two slices have the same number of rows.
	if len(slice1) != len(slice2) {
		return false
	}

	// Now, compare each row.
	for i := 0; i < len(slice1); i++ {
		// Check if the rows have the same length.
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}

		// Compare the elements of the rows.
		for j := 0; j < len(slice1[i]); j++ {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	solution_4 := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	solution_5 := [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}

	ret_4 := solution(4)
	fmt.Println("Solution 4 =", ret_4)
	if !Compare2DSlices(ret_4, solution_4) {
		fmt.Println("Solution 4 not match =", ret_4)
	}
	ret_5 := solution(5)
	fmt.Println("Solution 5 =", ret_5)
	if !Compare2DSlices(ret_5, solution_5) {
		fmt.Println("Solution 5 not match =", ret_5)
	}
}
