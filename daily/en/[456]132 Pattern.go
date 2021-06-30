//Given an array of n integers nums, a 132 pattern is a subsequence of three int
//egers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < n
//ums[j].
//
// Return true if there is a 132 pattern in nums, otherwise, return false.
//
// Follow up: The O(n^2) is trivial, could you come up with the O(n logn) or the
// O(n) solution?
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,4]
//Output: false
//Explanation: There is no 132 pattern in the sequence.
//
//
// Example 2:
//
//
//Input: nums = [3,1,4,2]
//Output: true
//Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
//
//
// Example 3:
//
//
//Input: nums = [-1,3,2,0]
//Output: true
//Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3,
// 0] and [-1, 2, 0].
//
//
//
// Constraints:
//
//
// n == nums.length
// 1 <= n <= 104
// -109 <= nums[i] <= 109
//
// Related Topics Stack
// 👍 2194 👎 139

package en

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/*
执行用时：60 ms, 在所有 Go 提交中击败了19.75%的用户
内存消耗：5.4 MB, 在所有 Go 提交中击败了11.11%的用户
*/
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	rans := make([][2]int, 0) // list of min--max
	rans = append(rans, [2]int{nums[0], nums[0]})
	for i := 1; i < len(nums); i++ {
		for _, ran := range rans {
			if nums[i] > ran[0] && nums[i] < ran[1] {
				return true
			}
		}

		lastRan := rans[len(rans)-1]
		if nums[i] > lastRan[1] {
			rans[len(rans)-1][1] = nums[i]
			continue
		}
		if nums[i] < lastRan[1] {
			rans = append(rans, [2]int{nums[i], nums[i]})
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

/*
[1,2,3,4]
[3,1,4,2]
[-1,3,2,0]
[2,4,5,1,3]
[1,3,2,4,5,6,7,8,9,10]

false
true
true
true
true

*/
// Monotonic stack
func find132pattern1(nums []int) bool {
	monotonicStack := newStack()
	monotonicStack.push(nums[len(nums)-1])
	var lastMin = math.MaxInt64

	for i := 0; i < len(nums); i++ {
		length := monotonicStack.len()
		min := monotonicStack.peek()
		top := min
		curr := nums[i]

		for monotonicStack.len() > 0 && curr > top {
			monotonicStack.pop()
			if monotonicStack.len() > 0 {
				top = monotonicStack.peek()
			}
		}
		if monotonicStack.len() == 0 && min < lastMin {
			lastMin = min
		}
		if curr > lastMin && monotonicStack.len() > 0 && monotonicStack.len() < length {
			return true
		}
		monotonicStack.push(curr)
	}

	return false
}

type stack []int

func newStack() (s stack) {
	s = make([]int, 0)
	return
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() (i int) {
	i = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) len() int {
	return len(*s)
}
