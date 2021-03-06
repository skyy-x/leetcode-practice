//Given a 32-bit signed integer, reverse digits of an integer. 
//
// Example 1: 
//
// 
//Input: 123
//Output: 321
// 
//
// Example 2: 
//
// 
//Input: -123
//Output: -321
// 
//
// Example 3: 
//
// 
//Input: 120
//Output: 21
// 
//
// Note: 
//Assume we are dealing with an environment which could only store integers with
//in the 32-bit signed integer range: [β231, 231 β 1]. For the purpose of this pro
//blem, assume that your function returns 0 when the reversed integer overflows. 
// Related Topics Math 
// π 3464 π 5468

package en

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
// θ§£η­ζε: ζ§θ‘θζΆ:0 ms,ε»θ΄₯δΊ100.00% ηGoη¨ζ· εε­ζΆθ:2.2 MB,ε»θ΄₯δΊ70.87% ηGoη¨ζ·
func reverse(x int) int {

	remainder:=1
	ret:=0
    for x!=0{
		remainder=x%10
		x/=10

		ret=ret*10+remainder
	}

	if ret>math.MaxInt32||ret<math.MinInt32{
		return 0
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
