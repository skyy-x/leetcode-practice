//Given a string, find the length of the longest substring without repeating cha
//racters. 
//
// 
// Example 1: 
//
// 
//Input: "abcabcbb"
//Output: 3 
//Explanation: The answer is "abc", with the length of 3. 
// 
//
// 
// Example 2: 
//
// 
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
// 
//
// 
// Example 3: 
//
// 
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3. 
//             Note that the answer must be a substring, "pwke" is a subsequence
// and not a substring.
// 
// 
// "umvejcuuk" => 6
// "aab" => 2
// 
// Related Topics Hash Table Two Pointers String Sliding Window 
// ð 9708 ð 586

package en

//leetcode submit region begin(Prohibit modification and deletion)
// è§£ç­æå: æ§è¡èæ¶:4 ms,å»è´¥äº94.25% çGoç¨æ· åå­æ¶è:3.2 MB,å»è´¥äº34.82% çGoç¨æ·
func lengthOfLongestSubstring(s string) int {
	dicMap:=make(map[int32]int)
	left:=0
	max:=0
	for idx,r:=range s{
		if i,ok:=dicMap[r];ok&&left<i+1{
			left=i+1
		}
		dicMap[r]=idx
		l:=idx-left+1
		if max<l{
			max=l
		}
	}
	return max
}
//leetcode submit region end(Prohibit modification and deletion)

/*
// pass è§£ç­æå: æ§è¡èæ¶:4 ms,å»è´¥äº32% çGoç¨æ· åå­æ¶è:3.2 MB,å»è´¥äº65% çGoç¨æ·
func lengthOfLongestSubstring(s string) int {
	dicMap:=make(map[int32]int)
	maxLen:=0
	leftIdx:=0
	l:=0
    for idx,r:=range s{
		if i,ok:=dicMap[r];!ok{
			l++
		}else{

			for ;leftIdx<i;leftIdx++{
				delete(dicMap,int32(s[leftIdx]))
			}
			l=idx-leftIdx
			leftIdx++
		}
		dicMap[r]=idx

		if l>maxLen{
			maxLen=l
		}
	}
    return maxLen
}
*/