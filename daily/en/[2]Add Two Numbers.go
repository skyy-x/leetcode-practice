//You are given two non-empty linked lists representing two non-negative integer
//s. The digits are stored in reverse order and each of their nodes contain a sing
//le digit. Add the two numbers and return it as a linked list. 
//
// You may assume the two numbers do not contain any leading zero, except the nu
//mber 0 itself. 
//
// Example: 
//
// 
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
// 
// Related Topics Linked List Math 
// 👍 8481 👎 2153
package en

 //Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }


//leetcode submit region begin(Prohibit modification and deletion)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil && l2==nil{
		return nil
	}

	carry:=0
	ret:=&ListNode{}
	current:=ret
	for l1!=nil || l2!=nil{
		if l1!=nil{
			carry+=l1.Val
			l1=l1.Next
		}
		if l2!=nil {
			carry += l2.Val
			l2 = l2.Next
		}

		current.Next=&ListNode{
			Val:  carry%10,
		}
		carry=carry/10
		current=current.Next
	}

	if carry>0{
		current.Next=&ListNode{Val:carry}
	}

	return ret.Next
}
//leetcode submit region end(Prohibit modification and deletion)
