package double_week_50

// 5717. 最少操作使数组递增
func minOperations(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	max := nums[0]
	adds := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= max {
			adds += max - nums[i] + 1
			max = max + 1
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return adds
}

// 5718. 统计一个圆中点的数目
func countPoints(points [][]int, queries [][]int) []int {
	pointMap := make(map[int]map[int]int) // x,y,cnt
	for _, p := range points {
		x := p[0]
		y := p[1]
		yMap, ok := pointMap[x]
		if !ok {
			yMap = make(map[int]int)
			pointMap[x] = yMap
		}
		yMap[y] += 1
	}

	ret := make([]int, len(queries))
	for i, q := range queries {
		ret[i] = cycleCnt(q, pointMap)
	}

	return ret
}

func cycleCnt(cycle []int, pointMap map[int]map[int]int) int {
	cx := cycle[0]
	cy := cycle[1]
	r := cycle[2]
	rr := r * r

	sum := 0
	for x, yMap := range pointMap {
		if x < cx-r || x > cx+r {
			continue
		}

		for y, cnt := range yMap {
			if y < cy-r || y > cy+r {
				continue
			}

			l := (x-cx)*(x-cx) + (y-cy)*(y-cy)
			if l <= rr {
				sum += cnt
			}
		}
	}

	return sum
}

// 5719. 每个查询的最大异或值
func getMaximumXor(nums []int, maximumBit int) []int {
	max := 0
	for i := 0; i < maximumBit; i++ {
		max = max | 1<<i
	}

	ret := make([]int, len(nums))
	x := 0
	for i := 0; i < len(nums); i++ {

		x = x ^ nums[i]
		ret[len(ret)-i-1] = (x ^ max) & max
	}
	return ret
}
