package e235

import (
	"sort"
	"strings"
)

// 1
// 5722. 截断句子
func truncateSentence(s string, k int) string {
	ss := strings.Split(s, " ")
	if len(ss) <= k {
		return s
	}
	return strings.Join(ss[:k], " ")
}

// 2
// 5723. 查找用户活跃分钟数
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	type userInfos struct {
		count int
		times map[int]struct{}
	}

	logMap := make(map[int]*userInfos)
	for _, userTime := range logs {
		id := userTime[0]
		time := userTime[1]
		info, ok := logMap[id]
		if !ok {
			info = &userInfos{
				count: 0,
				times: make(map[int]struct{}),
			}
			logMap[id] = info
		}

		_, ok = info.times[time]
		if !ok {
			info.times[time] = struct{}{}
			info.count++
		}
	}

	res := make([]int, k+1)
	for _, info := range logMap {
		if info.count > k {
			continue
		}
		res[info.count] += 1
	}
	return res[1:]
}

// 3
// 5724. 绝对差值和

// 使用二分查找优化 ac
type IndexValue struct {
	i int
	v int
}

type IndexValues []IndexValue

func (iv *IndexValues) Len() int {
	return len(*iv)
}

func (iv *IndexValues) Swap(i, j int) {
	(*iv)[i], (*iv)[j] = (*iv)[j], (*iv)[i]
}

func (iv *IndexValues) Less(i, j int) bool {
	return (*iv)[i].v > (*iv)[j].v
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {

	length := len(nums1)
	absoluteList := make(IndexValues, length)
	sum := 0
	for i := 0; i < length; i++ {
		abs := absolute(nums1[i], nums2[i])
		absoluteList[i] = IndexValue{
			i: i,
			v: abs,
		}
		sum += abs
	}
	sort.Sort(&absoluteList)
	sort.Ints(nums1)

	lastIncome := 0
	for _, iv := range absoluteList {

		income := maxIncome(nums1, nums2, iv)
		if income > lastIncome {
			lastIncome = income
		}
		if lastIncome >= iv.v {
			break
		}
	}
	return (sum - lastIncome) % (1e9 + 7)
}

func maxIncome(nums1, nums2 []int, iv IndexValue) int {

	expect := nums2[iv.i]

	left := 0
	right := len(nums1) - 1
	for right-left > 1 {
		mid := (right + left) / 2
		if nums1[mid] == expect {
			left, right = mid, mid
			break
		}
		if nums1[mid] < expect {
			left = mid
		}
		if nums1[mid] > expect {
			right = mid
		}
	}

	if left == right {
		return iv.v
	}

	abs := 0
	absl := absolute(expect, nums1[left])
	absr := absolute(expect, nums1[right])
	if absl < absr {
		abs = absl
	} else {
		abs = absr
	}

	return iv.v - abs
}

func absolute(i, j int) int {
	abs := i - j
	if abs < 0 {
		return -abs
	}
	return abs
}

// 4
