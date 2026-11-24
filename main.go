package main

import "fmt"

// UniquePaths counts the lattice paths from the top-left to the bottom-right
// of an m×n grid moving only right or down. Every path is a shuffle of
// (m-1) downs and (n-1) rights, so the answer is C(m+n-2, m-1), computed
// incrementally so intermediate values stay integral and in range.
func UniquePaths(m, n int) int64 {
	if m <= 0 || n <= 0 {
		return 0
	}
	k := int64(m - 1)
	total := int64(m + n - 2)
	res := int64(1)
	for i := int64(1); i <= k; i++ {
		res = res * (total - k + i) / i
	}
	return res
}

func main() {
	fmt.Println(UniquePaths(3, 7))
}
