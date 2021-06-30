package topic239

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
