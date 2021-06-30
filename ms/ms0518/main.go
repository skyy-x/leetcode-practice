package main

import "fmt"

// ac
// list [2,5,6,0,0,1,2]
func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 6))
	//fmt.Println(search([]int{2, 5, 6, 7, 8, 0, 1, 2}, 8))
	fmt.Println(search([]int{8, 8, 8, 8}, 8))
}

// 旋转数组中查找目标数
func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2

		if nums[mid] == target {
			return true
		}

		if nums[mid] >= nums[left] {
			// 左连续区间
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右连续区间
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
