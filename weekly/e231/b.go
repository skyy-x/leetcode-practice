package e231

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	c := goal - sum
	switch {
	case c > 0:
		return (c + limit - 1) / limit
	case c < 0:
		return (limit - 1 - c) / limit
	}
	return 0
}
