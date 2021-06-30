package main

import "fmt"

// ac
/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
For example:
Given the below binary tree and sum = 22,
          5
         / \
        4   8
       /   / \
      11  13  4
     /  \    / \
    7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]
*/
func main() {

	n := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Value: 7,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Value: 2,
				},
				Value: 11,
			},
			Right: &TreeNode{
				Left:  nil,
				Right: nil,
				Value: 0,
			},
			Value: 4,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Left:  nil,
				Right: nil,
				Value: 13,
			},
			Right: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: nil,
					Value: 5,
				},
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Value: 1,
				},
				Value: 4,
			},
			Value: 8,
		},
		Value: 5,
	}

	assist := NewAssist(22)
	assist.search(n, 0, make([]int, 0))
	fmt.Println(assist.results)
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

type Assist struct {
	targetSum int
	results   [][]int
}

func NewAssist(targetSum int) *Assist {
	return &Assist{
		results:   make([][]int, 0),
		targetSum: targetSum,
	}
}

func (a *Assist) search(n *TreeNode, sum int, stack []int) {
	sum = sum + n.Value

	if n.Left == nil && n.Right == nil && sum == a.targetSum {
		a.results = append(a.results, append(stack, n.Value))
	}

	if n.Left != nil {
		a.search(n.Left, sum, append(stack, n.Value))
	}

	if n.Right != nil {
		a.search(n.Right, sum, append(stack, n.Value))
	}
}
