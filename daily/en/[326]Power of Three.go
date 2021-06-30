//Given an integer, write a function to determine if it is a power of three. 
//
// Example 1: 
//
// 
//Input: 27
//Output: true
// 
//
// Example 2: 
//
// 
//Input: 0
//Output: false 
//
// Example 3: 
//
// 
//Input: 9
//Output: true 
//
// Example 4: 
//
// 
//Input: 45
//Output: false 
//
// Follow up: 
//Could you do it without using any loop / recursion? Related Topics Math 
// 👍 521 👎 1470

package en
//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	if n<1{
		return false
	}

    for n>1{
    	if n%3!=0{
    		return false
		}
    	n=n/3
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
