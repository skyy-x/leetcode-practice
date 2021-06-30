//Your country has an infinite number of lakes. Initially, all the lakes are emp
//ty, but when it rains over the nth lake, the nth lake becomes full of water. If
//it rains over a lake which is full of water, there will be a flood. Your goal is
// to avoid the flood in any lake.
//
// 👍 514 👎 110
// Given an integer array rains where:
//
//
// rains[i] > 0 means there will be rains over the rains[i] lake.
// rains[i] == 0 means there are no rains this day and you can choose one lake t
//his day and dry it.
//
//
// Return an array ans where:
//
//
// ans.length == rains.length
// ans[i] == -1 if rains[i] > 0.
// ans[i] is the lake you choose to dry in the ith day if rains[i] == 0.
//
//
// If there are multiple valid answers return any of them. If it is impossible t
//o avoid flood return an empty array.
//
// Notice that if you chose to dry a full lake, it becomes empty, but if you cho
//se to dry an empty lake, nothing changes. (see example 4)
//
//
// Example 1:
//
//
//Input: rains = [1,2,3,4]
//Output: [-1,-1,-1,-1]
//Explanation: After the first day full lakes are [1]
//After the second day full lakes are [1,2]
//After the third day full lakes are [1,2,3]
//After the fourth day full lakes are [1,2,3,4]
//There's no day to dry any lake and there is no flood in any lake.
//
//
// Example 2:
//
//
//Input: rains = [1,2,0,0,2,1]
//Output: [-1,-1,2,1,-1,-1]
//Explanation: After the first day full lakes are [1]
//After the second day full lakes are [1,2]
//After the third day, we dry lake 2. Full lakes are [1]
//After the fourth day, we dry lake 1. There is no full lakes.
//After the fifth day, full lakes are [2].
//After the sixth day, full lakes are [1,2].
//It is easy that this scenario is flood-free. [-1,-1,1,2,-1,-1] is another acce
//ptable scenario.
//
//
// Example 3:
//
//
//Input: rains = [1,2,0,1,2]
//Output: []
//Explanation: After the second day, full lakes are  [1,2]. We have to dry one l
//ake in the third day.
//After that, it will rain over lakes [1,2]. It's easy to prove that no matter w
//hich lake you choose to dry in the 3rd day, the other one will flood.
//
//
// Example 4:
//
//
//Input: rains = [69,0,0,0,69]
//Output: [-1,69,1,1,-1]
//Explanation: Any solution on one of the forms [-1,69,x,y,-1], [-1,x,69,y,-1] o
//r [-1,x,y,69,-1] is acceptable where 1 <= x,y <= 10^9
//
//
// Example 5:
//
//
//Input: rains = [10,20,20]
//Output: []
//Explanation: It will rain over lake 20 two consecutive days. There is no chanc
//e to dry any lake.
//
//
//
// Constraints:
//
//
// 1 <= rains.length <= 105
// 0 <= rains[i] <= 109
//
// Related Topics Array Hash Table
/*
解答成功: 执行耗时:152 ms,击败了100.00% 的Go用户 内存消耗:10.4 MB,击败了80.00% 的Go用户
*/
package en

//leetcode submit region begin(Prohibit modification and deletion)
func avoidFlood(rains []int) []int {
	ret := make([]int, len(rains))
	for i := range ret {
		ret[i] = -1
	}
	regionRains := make(map[int]int) // region => rainy day
	cl := make(clearDayList, 0)      // clear day
	for toDay, region := range rains {
		if region == 0 {
			cl.append(toDay)
			continue
		}

		lastRainDay, ok := regionRains[region]
		if !ok {
			regionRains[region] = toDay
			continue
		}

		// clear day is empty
		clearDay, ok := cl.pop(lastRainDay)
		if !ok {
			return []int{}
		}

		ret[clearDay] = region
		regionRains[region] = toDay
	}

	for len(cl) > 0 {
		ret[cl[0]] = 1
		cl = cl[1:]
	}
	return ret
}

type clearDayList []int

func (c *clearDayList) append(day int) {
	*c = append(*c, day)
}

func (c *clearDayList) pop(lastRainyDay int) (int, bool) {
	l := *c

	popIndex := -1
	day := -1
	for i, clearDay := range *c {
		if clearDay > lastRainyDay {
			popIndex = i
			day = clearDay
			break
		}
	}
	if popIndex == -1 {
		return day, false
	}

	if popIndex == 0 {
		*c = l[1:]
	} else if popIndex == len(l)-1 {
		*c = l[:len(l)-1]
	} else {
		*c = append(l[:popIndex], l[popIndex+1:]...)
	}
	return day, true
}

//leetcode submit region end(Prohibit modification and deletion)
