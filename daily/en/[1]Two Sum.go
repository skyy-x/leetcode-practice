//Given an array of integers, return indices of the two numbers such that they a
//dd up to a specific target. 
//
// You may assume that each input would have exactly one solution, and you may n
//ot use the same element twice. 
//
// Example: 
//
// 
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
// 
// Related Topics Array Hash Table 
// 👍 15763 👎 576

package en
//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	dictionary:=make(map[int]int)
	for idx,num:= range nums{
		if i,ok:=dictionary[target-num];ok{
			return []int{i,idx}
		}

		if _,ok:=dictionary[num];!ok{
			dictionary[num]=idx
		}
	}

	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
