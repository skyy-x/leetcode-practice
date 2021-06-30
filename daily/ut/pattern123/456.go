package pattern123

import "math"

// lc #456

// Monotonic stack
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	monotonicStack := newStack()

	mins := make([]int, len(nums))
	min := nums[0]
	for i, n := range nums {
		if min > n {
			mins[i] = n
			min = n
			continue
		}

		mins[i] = min
	}

	for i := len(nums) - 1; i > 0; i-- {
		curr := nums[i]

		top := math.MinInt64
		for monotonicStack.len() > 0 && curr > monotonicStack.peek() {
			top = monotonicStack.pop()
		}

		if top > mins[i-1] {
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
