package e233

import (
	"container/heap"
)

// 1
// 5709. 最大升序子数组和
func maxAscendingSum(nums []int) int {
	maxSum := 0
	currSum := 0
	prev := 0
	for _, num := range nums {
		if num <= prev {
			prev = num
			currSum = num
			continue
		}
		if num > prev {
			currSum += num
		}
		if currSum > maxSum {
			maxSum = currSum
		}
		prev = num
	}
	return maxSum
}

type buyHeap []*order
type sellHeap []*order

type order struct {
	price  int
	amount int
}

// Len is the number of elements in the collection.
func (h *buyHeap) Len() int {
	return len(*h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (h *buyHeap) Less(i, j int) bool {
	return (*h)[i].price > (*h)[j].price
}

// Swap swaps the elements with indexes i and j.
func (h *buyHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push add x as element Len()
func (h *buyHeap) Push(x interface{}) {
	*h = append(*h, x.(*order))
}

// Pop remove and return element Len() - 1.
func (h *buyHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *buyHeap) PeekOrder() *order {
	if len(*h) == 0 {
		return nil
	}
	return (*h)[0]
}

// Len is the number of elements in the collection.
func (h *sellHeap) Len() int {
	return len(*h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (h *sellHeap) Less(i, j int) bool {
	return (*h)[i].price < (*h)[j].price
}

// Swap swaps the elements with indexes i and j.
func (h *sellHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push add x as element Len()
func (h *sellHeap) Push(x interface{}) {
	*h = append(*h, x.(*order))
}

// Pop remove and return element Len() - 1.
func (h *sellHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *sellHeap) PeekOrder() *order {
	if len(*h) == 0 {
		return nil
	}
	return (*h)[0]
}

// 2
// 5710. 积压订单中的订单总数
func getNumberOfBacklogOrders(orders [][]int) int {

	buyOrder := &buyHeap{}
	sellOrder := &sellHeap{}

	for _, o := range orders {
		// buy
		if o[2] == 0 {
			buy(sellOrder, buyOrder, o[0], o[1])
		}
		// sell
		if o[2] == 1 {
			sell(sellOrder, buyOrder, o[0], o[1])
		}
	}

	var total int64
	for _, o := range *buyOrder {
		total += int64(o.amount)
	}
	for _, o := range *sellOrder {
		total += int64(o.amount)
	}

	return int(total % (1e9 + 7))
}

func buy(sellOrder *sellHeap, buyOrder *buyHeap, price, amount int) {
	if amount == 0 {
		return
	}
	s := sellOrder.PeekOrder()
	if s == nil || s.price > price {
		heap.Push(buyOrder, &order{
			price:  price,
			amount: amount,
		})
		return
	}

	if s.amount > amount {
		s.amount -= amount
		return
	}
	heap.Pop(sellOrder)
	buy(sellOrder, buyOrder, price, amount-s.amount)
}

func sell(sellOrder *sellHeap, buyOrder *buyHeap, price, amount int) {
	if amount == 0 {
		return
	}
	s := buyOrder.PeekOrder()
	if s == nil || s.price < price {
		heap.Push(sellOrder, &order{
			price:  price,
			amount: amount,
		})
		return
	}

	if s.amount > amount {
		s.amount -= amount
		return
	}
	heap.Pop(buyOrder)
	sell(sellOrder, buyOrder, price, amount-s.amount)
}

// 3
// 5696. 统计异或值在范围内的数对有多少
/*
给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数：low 和 high ，请返回 漂亮数对 的数目。
漂亮数对 是一个形如 (i, j) 的数对，其中 0 <= i < j < nums.length 且 low <= (nums[i] XOR nums[j]) <= high 。


示例 1：

输入：nums = [1,4,2,7], low = 2, high = 6
输出：6
解释：所有漂亮数对 (i, j) 列出如下：
    - (0, 1): nums[0] XOR nums[1] = 5
    - (0, 2): nums[0] XOR nums[2] = 3
    - (0, 3): nums[0] XOR nums[3] = 6
    - (1, 2): nums[1] XOR nums[2] = 6
    - (1, 3): nums[1] XOR nums[3] = 3
    - (2, 3): nums[2] XOR nums[3] = 5
示例 2：

输入：nums = [9,8,4,2,1], low = 5, high = 14
输出：8
解释：所有漂亮数对 (i, j) 列出如下：
 - (0, 2): nums[0] XOR nums[2] = 13
 - (0, 3): nums[0] XOR nums[3] = 11
 - (0, 4): nums[0] XOR nums[4] = 8
 - (1, 2): nums[1] XOR nums[2] = 12
 - (1, 3): nums[1] XOR nums[3] = 10
 - (1, 4): nums[1] XOR nums[4] = 9
 - (2, 3): nums[2] XOR nums[3] = 6
 - (2, 4): nums[2] XOR nums[4] = 5

提示：

1 <= nums.length <= 2 * 1e4
1 <= nums[i] <= 2 * 1e4
1 <= low <= high <= 2 * 1e4

*/
func maxValue(n int, index int, maxSum int) int {
	sum := n
	layer := 1
	for sum <= maxSum {
		left := layer - 1
		right := layer - 1
		if left > index {
			left = index
		}
		if right > n-index-1 {
			right = n - index - 1
		}
		if left < layer-1 && right < layer-1 {
			layer++
			break
		}
		sum = sum + left + right + 1
		layer++
	}
	if maxSum > sum {
		layer += (maxSum - sum) / n
	}
	return layer - 1
}

// 4
func countPairs(nums []int, low int, high int) int {
	numTimes := make(map[int]int)
	uniqueNums := make([]int, 0)
	for _, n := range nums {
		t, ok := numTimes[n]
		if !ok {
			uniqueNums = append(uniqueNums, n)
		}
		numTimes[n] = t + 1
	}

	ret := 0
	for i := 0; i < len(uniqueNums); i++ {
		for j := i + 1; j < len(uniqueNums); j++ {
			xor := uniqueNums[i] ^ uniqueNums[j]
			if xor >= low && xor <= high {
				ret += numTimes[uniqueNums[i]] * numTimes[uniqueNums[j]]
			}
		}
	}
	return ret
}

// trie树
// 通过	144 ms	21.6 MB
const length = 2e4 + 10

func countPairs1(nums []int, low int, high int) int {
	t := newTrie()
	count := 0
	for _, n := range nums {
		count += t.search(n, high) - t.search(n, low-1)
		t.insert(n)
	}
	return count
}

type trie struct {
	// 0为最低位
	nodes [length * 20][2]int // node index => 0/1 => next node index
	count [length * 15]int    // node count
	index int
}

func newTrie() *trie {
	return &trie{}
}

// 动态增加节点索引
func (t *trie) insert(n int) {
	nodeIndex := 0
	for i := 15; i > -1; i-- {
		bit := n >> i & 1
		if t.nodes[nodeIndex][bit] == 0 {
			t.index++
			t.nodes[nodeIndex][bit] = t.index
		}
		// 下一个节点索引
		nodeIndex = t.nodes[nodeIndex][bit]
		// 记录节点出现次数
		t.count[nodeIndex]++
	}
}

// 查出n与所有异或值<=max的个数
func (t *trie) search(n, max int) (pairs int) {
	nodeIndex := 0
	for i := 15; i > -1; i-- {
		xor := max >> i & 1
		bit := n >> i & 1
		if xor == 0 {
			// 异或值为0，向bit一样的走，继续计算
			sameIndex := t.nodes[nodeIndex][bit]
			if sameIndex == 0 {
				return
			}
			nodeIndex = sameIndex
		} else {
			// 异或值为1
			// bit一样的全符合
			if sameIndex := t.nodes[nodeIndex][bit]; sameIndex != 0 {
				pairs += t.count[nodeIndex]
			}
			// 不一样的进行下次循环
			diffIndex := t.nodes[nodeIndex][1^bit]
			if diffIndex == 0 {
				// 没有不一样的，即所有高位xor为0<1
				return
			}
			// 有不一样的，向下走
			nodeIndex = diffIndex
		}
	}
	// 结尾一个位，无论异或值为0/1需要计算
	pairs += t.count[nodeIndex]
	return
}
