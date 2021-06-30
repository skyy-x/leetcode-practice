package e232

import (
	"container/heap"
)

// 5701. 仅执行一次字符串交换能否使两个字符串相等
// 给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。
// 如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。
/*
示例 1：

输入：s1 = "bank", s2 = "kanb"
输出：true
解释：例如，交换 s2 中的第一个和最后一个字符可以得到 "bank"
示例 2：

输入：s1 = "attack", s2 = "defend"
输出：false
解释：一次字符串交换无法使两个字符串相等
示例 3：

输入：s1 = "kelb", s2 = "kelb"
输出：true
解释：两个字符串已经相等，所以不需要进行字符串交换
示例 4：

输入：s1 = "abcd", s2 = "dcba"
输出：false
*/
//
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	ok := false
	length := len(s1)
	lastIdx := -1
	for i := 0; i < length; i++ {
		if s1[i] == s2[i] {
			continue
		}

		if lastIdx == -1 {
			lastIdx = i
			continue
		}
		if ok {
			return false
		}

		if s1[lastIdx] == s2[i] && s1[i] == s2[lastIdx] {
			ok = true
		} else {
			return false
		}

	}
	return ok
}

// 5702. 找出星型图的中心节点
func findCenter(edges [][]int) int {
	numsMap := make(map[int]struct{})
	for i := 0; i < len(edges); i++ {
		if _, ok := numsMap[edges[i][0]]; ok {
			return edges[i][0]
		}
		numsMap[edges[i][0]] = struct{}{}
		if _, ok := numsMap[edges[i][1]]; ok {
			return edges[i][1]
		}
		numsMap[edges[i][1]] = struct{}{}
	}
	return 0
}

// 5703. 最大平均通过率
// 一所学校里有一些班级，每个班级里有一些学生，现在每个班都会进行一场期末考试。给你一个二维数组 classes ，其中 classes[i] = [passi, totali] ，表示你提前知道了第 i 个班级总共有 totali 个学生，其中只有 passi 个学生可以通过考试。
// 给你一个整数 extraStudents ，表示额外有 extraStudents 个聪明的学生，他们 一定 能通过任何班级的期末考。你需要给这 extraStudents 个学生每人都安排一个班级，使得 所有 班级的 平均 通过率 最大 。
// 一个班级的 通过率 等于这个班级通过考试的学生人数除以这个班级的总人数。平均通过率 是所有班级的通过率之和除以班级数目。
// 请你返回在安排这 extraStudents 个学生去对应班级后的 最大 平均通过率。与标准答案误差范围在 10-5 以内的结果都会视为正确结果。
/*
示例 1：

输入：classes = [[1,2],[3,5],[2,2]], extraStudents = 2
输出：0.78333
解释：你可以将额外的两个学生都安排到第一个班级，平均通过率为 (3/4 + 3/5 + 2/2) / 3 = 0.78333 。
示例 2：

输入：classes = [[2,4],[3,9],[4,5],[2,10]], extraStudents = 4
输出：0.53485


提示：

1 <= classes.length <= 105
classes[i].length == 2
1 <= passi <= totali <= 105
1 <= extraStudents <= 105
*/

//
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	c := make(Classes, len(classes))
	for i, class := range classes {
		c[i] = Item{
			i:     i,
			minus: float64(class[1]-class[0]) / float64(class[1]*(class[1]+1)),
		}
	}
	heap.Init(&c)
	for ; extraStudents > 0; extraStudents-- {
		item := heap.Pop(&c).(Item)
		class := classes[item.i]
		class[0] += 1
		class[1] += 1
		item.minus = float64(class[1]-class[0]) / float64(class[1]*(class[1]+1))
		heap.Push(&c, item)
	}
	sum := 0.0000000
	for _, class := range classes {
		sum += float64(class[0]) / float64(class[1])
	}
	return sum / float64(c.Len())
}

type Classes []Item

func (p *Classes) Len() int { return len(*p) }
func (p *Classes) Less(i, j int) bool {

	return (*p)[j].minus < (*p)[i].minus
}
func (p *Classes) Swap(i, j int) { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

type Item struct {
	i     int
	minus float64
}

func (p *Classes) Push(i interface{}) {
	item := i.(Item)
	*p = append(*p, item)
}

func (p *Classes) Pop() interface{} {
	i := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]

	return i
}

// 5704. 好子数组的最大分数
/*
给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。

一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个 好 子数组的两个端点下标需要满足 i <= k <= j 。

请你返回 好 子数组的最大可能 分数 。



示例 1：

输入：nums = [1,4,3,7,4,5], k = 3
输出：15
解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。
示例 2：

输入：nums = [5,5,4,5,4,1,1,1], k = 0
输出：20
解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。

bad case:
[6569,9667,3148,7698,1622,2194,793,9041,1670,1872] 5
9732


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 2 * 104
0 <= k < nums.length
*/
// 我的思路：双指针+剪枝
func maximumScore(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if k > len(nums)-1 {
		return 0
	}

	left := k
	right := k
	min := nums[k]
	max := min
	for {
		for left > 0 {
			if nums[left-1] >= min {
				left--
				continue
			}
			break
		}
		for right < len(nums)-1 {
			if nums[right+1] >= min {
				right++
				continue
			}
			break
		}
		good := min * (right - left + 1)
		if good > max {
			max = good
		}

		moveToLeft := left
		if left > 0 && nums[left-1]*(right) > good {
			moveToLeft--
		}
		moveToRight := right
		if right < len(nums)-1 && nums[right+1]*(len(nums)-left) > good {
			moveToRight++
		}

		/*
			moveToLeft := left
			if left-1 >= 0 {
				moveToLeft = left - 1
			}
			moveToRight := right
			if right+1 <= len(nums)-1 {
				moveToRight = right + 1
			}

		*/

		if left == moveToLeft && right == moveToRight {
			break
		}
		if moveToLeft < left && moveToRight > right {
			if nums[moveToLeft] == nums[moveToRight] {
				left = moveToLeft
				right = moveToRight
			} else if nums[moveToLeft] > nums[moveToRight] {
				left = moveToLeft
			} else {
				right = moveToRight
			}
		} else if moveToLeft < left {
			left = moveToLeft
		} else {
			right = moveToRight
		}

		min = Min(Min(nums[left], min), Min(nums[right], min))
	}

	return max
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
