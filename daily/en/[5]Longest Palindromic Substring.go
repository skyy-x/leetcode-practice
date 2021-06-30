//Given a string s, find the longest palindromic substring in s. You may assume 
//that the maximum length of s is 1000. 
//
// Example 1: 
//
// 
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
// 
//
// Example 2: 
//
// 
//Input: "cbbd"
//Output: "bb"
// 
// Related Topics String Dynamic Programming 
// 👍 7204 👎 545

package en
//leetcode submit region begin(Prohibit modification and deletion)
// 解答成功: 执行耗时:8 ms,击败了75.83% 的Go用户 内存消耗:2.2 MB,击败了100.00% 的Go用户
func longestPalindrome(s string) string {
	if len(s)==0{
		return ""
	}
    maxLength:=1
    maxL:=-1
    maxR:=1
	for idx,rune:=range s{
		w:=uint8(rune)
		l:=idx-1
		r:=idx+1
		for l>=0&&s[l]==w{
			l--
			continue
		}
		for r<len(s)&&s[r]==w{
			r++
			continue
		}
		for l>=0&&r<len(s){
			if s[l]!=s[r]{
				break
			}
			l--
			r++
		}
		if r-l-1>maxLength{
			maxLength=r-l-1
			maxL=l
			maxR=r
		}
	}

	return s[maxL+1:maxR]
}

//leetcode submit region end(Prohibit modification and deletion)
