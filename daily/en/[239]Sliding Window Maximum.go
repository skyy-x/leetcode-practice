//You are given an array of integers nums, there is a sliding window of size k w
//hich is moving from the very left of the array to the very right. You can only s
//ee the k numbers in the window. Each time the sliding window moves right by one
//position.
//
// Return the max sliding window.
//
//
// Example 1:
//
//
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
//Output: [3,3,5,5,6,7]
//Explanation:
//Window position                Max
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// Example 2:
//
//
//Input: nums = [1], k = 1
//Output: [1]
//
//
// Example 3:
//
//
//Input: nums = [1,-1], k = 1
//Output: [1,-1]
//
//
// Example 4:
//
//
//Input: nums = [9,11], k = 2
//Output: [11]
//
//
// Example 5:
//
//
//Input: nums = [4,-2], k = 2
//Output: [4]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics Heap Sliding Window Dequeue
// 👍 5509 👎 229

package en

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	s := newMonotonicStack(k)
	ret := make([]int, len(nums)-k+1)
	for i, n := range nums {
		s.append(element{
			value: n,
			index: i,
		})
		retIndex := i - k + 1
		if retIndex > -1 {
			ret[retIndex] = s.maxValue()
		}
	}
	return ret
}

type element struct {
	value int
	index int
}

type monotonicStack struct {
	// 大 -> 小
	elements []element
	capacity int
}

func newMonotonicStack(cap int) (s monotonicStack) {
	s = monotonicStack{
		elements: make([]element, 0),
		capacity: cap,
	}
	return
}

func (s *monotonicStack) append(e element) {
	if len(s.elements) == 0 {
		s.elements = append(s.elements, e)
		return
	}

	// 超长出队
	lastIndex := e.index - s.capacity
	if s.elements[0].index <= lastIndex {
		s.elements = s.elements[1:]
	}

	// 小的出栈
	for len(s.elements) > 0 && s.elements[len(s.elements)-1].value < e.value {
		s.elements = s.elements[:len(s.elements)-1]
	}
	s.elements = append(s.elements, e)
}

func (s *monotonicStack) maxValue() int {
	return s.elements[0].value
}

//leetcode submit region end(Prohibit modification and deletion)
