package main

import (
	"fmt"
	"sort"
)

// ac
// 给定数组 -1, 0, 1, 2, -1, -4，返回和为0且不重复的组合
func main() {
	//ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4, 0, 0, 0}))
}

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)
	//fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > 0 && j != i+1 && nums[j-1] == nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > 0 && k != j+1 && nums[k-1] == nums[k] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return res
}
