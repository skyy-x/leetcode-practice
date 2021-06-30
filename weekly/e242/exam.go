package e242

// 5763. 哪种连续子字符串更长
/*
给你一个二进制字符串 s 。如果字符串中由 1 组成的 最长 连续子字符串 严格长于 由 0 组成的 最长 连续子字符串，返回 true ；否则，返回 false 。

例如，s = "110100010" 中，由 1 组成的最长连续子字符串的长度是 2 ，由 0 组成的最长连续子字符串的长度是 3 。
注意，如果字符串中不存在 0 ，此时认为由 0 组成的最长连续子字符串的长度是 0 。字符串中不存在 1 的情况也适用此规则。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longer-contiguous-segments-of-ones-than-zeros
*/
func checkZeroOnes(s string) bool {
	var zeroCnt, oneCnt, zeroCntMax, oneCntMax int

	add := func(r int32) {
		if r == '1' {
			oneCnt++
			if oneCnt > oneCntMax {
				oneCntMax = oneCnt
			}
		} else {
			zeroCnt++
			if zeroCnt > zeroCntMax {
				zeroCntMax = zeroCnt
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if i > 0 && s[i-1] == '1' {
				add('1')
			} else {
				oneCnt = 0
				add('1')
			}
			continue
		}

		if i > 0 && s[i-1] == '0' {
			add('0')
		} else {
			zeroCnt = 0
			add('0')
		}
	}

	return oneCntMax > zeroCntMax
}

// 5765. 跳跃游戏 VII
/*
给你一个下标从 0 开始的二进制字符串 s 和两个整数 minJump 和 maxJump 。一开始，你在下标 0 处，且该位置的值一定为 '0' 。当同时满足如下条件时，你可以从下标 i 移动到下标 j 处：

i + minJump <= j <= min(i + maxJump, s.length - 1) 且
s[j] == '0'.
如果你可以到达 s 的下标 s.length - 1 处，请你返回 true ，否则返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-vii
*/
