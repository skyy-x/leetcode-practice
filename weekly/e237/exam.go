package e237

import (
	"container/heap"
	"fmt"
	"sort"
)

// 5734. 判断句子是否为全字母句
func checkIfPangram(sentence string) bool {
	wordsMap := make(map[int32]struct{})
	for _, w := range sentence {
		ascii := w - 'a'
		if ascii > -1 && ascii < 26 {
			wordsMap[ascii] = struct{}{}
		}
	}
	return len(wordsMap) == 26
}

// 5735. 雪糕的最大数量
func maxIceCream(costs []int, coins int) int {

	sort.Ints(costs)
	cnt := 0
	for _, c := range costs {
		coins -= c
		if coins < 0 {
			break
		}
		cnt++
	}
	return cnt
}

// 5736. 单线程 CPU
func getOrder(tasks [][]int) []int {
	pq := NewPQ()
	ret := make([]int, 0)
	taskList := newTasks(tasks)

	runTo := 0
	runs := make([]int, 0)
	for i := 0; i < len(taskList); i++ {
		t := taskList[i]
		if t.start <= runTo || pq.Len() == 0 {
			heap.Push(pq, &t)
			continue
		} else {
			i--
		}

		popT := heap.Pop(pq).(*task)
		ret = append(ret, popT.index)
		now := runTo
		if popT.start > now {
			now = popT.start
		}
		runTo = now + popT.duration
		runs = append(runs, runTo)
	}

	for pq.Len() > 0 {
		t := heap.Pop(pq).(*task)
		ret = append(ret, t.index)
	}
	fmt.Println(runs)
	fmt.Println(ret)
	return ret
}

type tasks []task

func newTasks(list [][]int) tasks {
	taskList := tasks(make([]task, len(list)))

	for i, t := range list {
		taskList[i] = task{
			start:    t[0],
			index:    i,
			duration: t[1],
		}
	}
	sort.Sort(taskList)
	return taskList
}

func (t tasks) Len() int {
	return len(t)
}

func (t tasks) Less(i, j int) bool {
	if t[i].start < t[j].start {
		return true
	}
	if t[i].start == t[j].start && t[i].duration < t[j].duration {
		return true
	}
	if t[i].start == t[j].start && t[i].duration == t[j].duration && t[i].index < t[j].index {
		return true
	}
	return false
}

func (t tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type task struct {
	start    int
	index    int
	duration int
}

type PQ []*task

func NewPQ() *PQ {
	pq := PQ(make([]*task, 0))
	return &pq
}

func (pq *PQ) Len() int { return len(*pq) }

func (pq *PQ) Less(i, j int) bool {
	if (*pq)[i].duration < (*pq)[j].duration {
		return true
	}
	if (*pq)[i].duration == (*pq)[j].duration && (*pq)[i].index < (*pq)[j].index {
		return true
	}
	return false
}

func (pq *PQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

// Push
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*task))
}

// Pop
func (pq *PQ) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}

// Peek
func (pq *PQ) Peek() *task {
	return (*pq)[len(*pq)-1]
}

// 5737. 所有数对按位与结果的异或和
func getXORSum(arr1 []int, arr2 []int) int {
	xorSum2 := 0
	xorSum := 0
	for _, n := range arr2 {
		xorSum2 ^= n
	}
	for _, n := range arr1 {
		xorSum ^= n & xorSum2
	}
	return xorSum
}
