package double_week_53

import (
	"fmt"
	"sort"
)

// 5754. 长度为三且各字符不同的子字符串
/*
如果一个字符串不含有任何重复字符，我们称这个字符串为 好 字符串。
给你一个字符串 s ，请你返回 s 中长度为 3 的 好子字符串 的数量。
注意，如果相同的好子字符串出现多次，每一次都应该被记入答案之中。
子字符串 是一个字符串中连续的字符序列。

示例 1：
输入：s = "xyzzaz"
输出：1
解释：总共有 4 个长度为 3 的子字符串："xyz"，"yzz"，"zza" 和 "zaz" 。
唯一的长度为 3 的好子字符串是 "xyz" 。

示例 2：
输入：s = "aababcabc"
输出：4
解释：总共有 7 个长度为 3 的子字符串："aab"，"aba"，"bab"，"abc"，"bca"，"cab" 和 "abc" 。
好子字符串包括 "abc"，"bca"，"cab" 和 "abc" 。
*/
func countGoodSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] || s[i+1] == s[i+2] || s[i] == s[i+2] {
			continue
		}
		fmt.Println(s[i], s[i+1], s[i+2])
		cnt++
	}
	return cnt
}

// 5755. 数组中最大数对和的最小值
func minPairSum(nums []int) int {

	max := 0
	sort.Ints(nums)
	l := len(nums) / 2
	for i := 0; i < l; i++ {
		sum := nums[i] + nums[len(nums)-i-1]
		if sum > max {
			max = sum
		}
	}
	return max
}

// 5757. 矩阵中最大的三个菱形和
func getBiggestThree(grid [][]int) []int {
	h := len(grid)
	if h == 0 {
		return []int{0, 0, 0}
	}
	w := len(grid[0])

	maxLen := (min(h, w) + 1) / 2
	res := make([]int, 0)
	exclude := make(map[int]struct{})

	for k := 0; k < maxLen; k++ {
		for i := k; i < h-k; i++ {
			for j := k; j < w-k; j++ {
				sum := 0
				if k > 0 {
					for x := 0; x < k; x++ {
						sum += grid[i+k-x][j+x]
						sum += grid[i-k+x][j-x]
						sum += grid[i-x][j+k-x]
						sum += grid[i+x][j-k+x]
					}
				} else {
					sum = grid[i][j]
				}
				if _, ok := exclude[sum]; ok {
					continue
				}
				res = append(res, sum)
				exclude[sum] = struct{}{}
			}
		}
	}

	sort.Sort(IntSlice(res))
	if len(res) <= 3 {
		return res
	}
	return res[:3]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] > p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
