package e243

import "strconv"

// 5772. 检查某单词是否等于两单词之和
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return value(firstWord)+value(secondWord) == value(targetWord)
}

func value(word string) int {
	var res int

	for i := 0; i < len(word); i++ {

		var e int = 1
		for j := 0; j < len(word)-i-1; j++ {
			e *= 10
		}
		res += int(word[i]-'a') * e
	}

	return res
}

// 5773. 插入后的最大值
func maxValue(n string, x int) string {
	var reverse bool
	if n[0] == '-' {
		reverse = true
	}

	var i int
	if reverse {
		i = 1
	}

	for ; i < len(n); i++ {

		if reverse {
			if n[i]-'0' > uint8(x) {
				return n[:i] + strconv.Itoa(x) + n[i:]
			}
		} else {
			if n[i]-'0' < uint8(x) {
				return n[:i] + strconv.Itoa(x) + n[i:]
			}
		}

	}

	return n + strconv.Itoa(x)
}
