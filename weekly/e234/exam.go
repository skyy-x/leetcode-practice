package e234

import "bytes"

// 1
func numDifferentIntegers(word string) int {
	nums := make(map[int32]struct{})
	var lastNum int32 = -1
	for _, w := range word {
		n := w - '0'
		if n > -1 && n < 10 {
			if lastNum == -1 {
				lastNum = n
				continue
			}
			lastNum = lastNum*10 + n
			continue
		}
		if lastNum > -1 {
			nums[lastNum] = struct{}{}
			lastNum = -1
		}
	}

	if lastNum > -1 {
		nums[lastNum] = struct{}{}
	}
	return len(nums)
}

// 2
func reinitializePermutation(n int) int {
	if n == 2 {
		return 1
	}
	dp := make(map[int]int) // 记录已有数字
	for i := 0; i < n; i++ {
		rule := i%2 == 0
		if rule {
			dp[i] = i / 2
		} else {
			dp[i] = n/2 + (i-1)/2
		}
	}

	idx := dp[1]
	cnt := 0
	for idx != 1 {
		idx = dp[idx]
		cnt++
	}
	return cnt + 1
}

// 3
func evaluate(s string, knowledge [][]string) string {
	kMap := make(map[string]string)
	for _, kv := range knowledge {
		kMap[kv[0]] = kv[1]
	}

	b := &bytes.Buffer{}
	lastKey := ""
	isKey := false
	for _, w := range s {
		if w == '(' {
			isKey = true
			continue
		}
		if !isKey {
			b.WriteRune(w)
		}
		if w == ')' {
			kWord, ok := kMap[lastKey]
			if ok {
				b.WriteString(kWord)
			} else {
				b.WriteString("?")
			}
			lastKey = ""
			isKey = false
			continue
		}
		if isKey {
			lastKey += string(w)
		}
	}
	return b.String()
}

// 4
