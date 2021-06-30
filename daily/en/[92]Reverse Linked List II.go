//Given the head of a singly linked list and two integers left and right where l
//eft <= right, reverse the nodes of the list from position left to position right
//, and return the reversed list.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], left = 2, right = 4
//Output: [1,4,3,2,5]
//
//
// Example 2:
//
//
//Input: head = [5], left = 1, right = 1
//Output: [5]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is n.
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//Follow up: Could you do it in one pass? Related Topics Linked List
// 👍 3431 👎 178

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package en

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了33.27%的用户
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}
	headParent := &ListNode{
		Val:  0,
		Next: head,
	}
	var prev, prefix, suffix *ListNode
	c := headParent
	cnt := 0
	for ; c != nil; cnt++ {
		next := c.Next

		if cnt < left {
			prev = c
			c = next
			continue
		} else if cnt == left {
			prefix = prev
			suffix = c

			prev = c
			c = next
			continue
		}

		c.Next = prev

		if cnt == right {
			suffix.Next = next
			prefix.Next = c
			break
		}

		prev = c
		c = next
	}
	return headParent.Next
}

//leetcode submit region end(Prohibit modification and deletion)
