package main

// ac
// bst
type Node struct {
	left  *Node
	right *Node
	value int
}

func main() {

}

type search struct {
	count int
	k     int
	ret   int
}

// 二叉搜索树第k大
func (s *search) dfs(n *Node) {

	if n.right != nil {
		s.dfs(n.right)
	}
	s.count++
	if s.count == s.k {
		s.ret = n.value
		return
	}
	if n.left != nil {
		s.dfs(n.left)
	}
}

func (s *search) dfs1(n *Node) {

	stack := make([]*Node, 0)
	stack = append(stack, n)

	for s.count < s.k {

		for thisN := stack[len(stack)-1]; thisN.right != nil; {
			thisN = thisN.right
			stack = append(stack, thisN)
		}

		if len(stack) == 0 {
			return
		}
		thisN := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		s.count++
		if s.count == s.k {
			s.ret = n.value
			return
		}

		stack = append(stack, thisN.left)
	}
}
