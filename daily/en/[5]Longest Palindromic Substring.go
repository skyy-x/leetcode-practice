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
// ð 7204 ð 545

package en
//leetcode submit region begin(Prohibit modification and deletion)
// è§£ç­æå: æ§è¡èæ¶:8 ms,å»è´¥äº75.83% çGoç¨æ· åå­æ¶è:2.2 MB,å»è´¥äº100.00% çGoç¨æ·
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
