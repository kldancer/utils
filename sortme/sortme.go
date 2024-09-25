package sortme

import "sort"

// 二分查找标准库
func me() {
	a := []int{1, 2, 3, 4}
	target := 4
	sort.SearchInts(a, target)
}

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
