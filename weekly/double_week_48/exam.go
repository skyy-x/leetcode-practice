package double_week_48

import "sort"

// 1
func secondHighest(s string) int {

	max := -1
	secondMax := -1
	for _, x := range s {
		if x < 48 || x > 57 {
			continue
		}

		n := int(x - 48)
		if n > max {
			secondMax = max
			max = n
		} else if n == max {
			continue
		} else if n > secondMax {
			secondMax = n
		}

	}

	return secondMax
}

// 2
type AuthenticationManager struct {
	ttl    int
	tokens map[string]int // Unexpired time
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:    timeToLive,
		tokens: make(map[string]int),
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	unexpiredTime, ok := this.tokens[tokenId]
	if !ok {
		return
	}

	if unexpiredTime > currentTime {
		this.tokens[tokenId] = currentTime + this.ttl
	} else {
		delete(this.tokens, tokenId)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	c := 0
	for _, unexpiredTime := range this.tokens {
		if unexpiredTime <= currentTime {
			continue
		}
		c++
	}
	return c
}

// 3
// 5712. 你能构造出连续值的最大数目
/*
给你一个长度为 n 的整数数组 coins ，它代表你拥有的 n 个硬币。第 i 个硬币的值为 coins[i] 。如果你从这些硬币中选出一部分硬币，它们的和为 x ，那么称，你可以 构造 出 x 。
请返回从 0 开始（包括 0 ），你最多能 构造 出多少个连续整数。
你可能有多个相同值的硬币。

示例 1：
输入：coins = [1,3]
输出：2
解释：你可以得到以下这些值：
- 0：什么都不取 []
- 1：取 [1]
从 0 开始，你可以构造出 2 个连续整数。

示例 2：
输入：coins = [1,1,1,4]
输出：8
解释：你可以得到以下这些值：
- 0：什么都不取 []
- 1：取 [1]
- 2：取 [1,1]
- 3：取 [1,1,1]
- 4：取 [4]
- 5：取 [4,1]
- 6：取 [4,1,1]
- 7：取 [4,1,1,1]
从 0 开始，你可以构造出 8 个连续整数。

示例 3：
输入：nums = [1,4,10,3,1]
输出：20

提示：
coins.length == n
1 <= n <= 4 * 104
1 <= coins[i] <= 4 * 104
*/
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	max := 0
	for _, c := range coins {
		if c > max+1 {
			break
		}
		max += c
	}
	return max + 1
}

// 4
// 5695. N 次操作后的最大分数和
/*
给你 nums ，它是一个大小为 2 * n 的正整数数组。你必须对这个数组执行 n 次操作。
在第 i 次操作时（操作编号从 1 开始），你需要：

选择两个元素 x 和 y 。
获得分数 i * gcd(x, y) 。
将 x 和 y 从 nums 中删除。
请你返回 n 次操作后你能获得的分数和最大为多少。

函数 gcd(x, y) 是 x 和 y 的最大公约数。

示例 1：
输入：nums = [1,2]
输出：1
解释：最优操作是：
(1 * gcd(1, 2)) = 1

示例 2：
输入：nums = [3,4,6,8]
输出：11
解释：最优操作是：
(1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11

示例 3：
输入：nums = [1,2,3,4,5,6]
输出：14
解释：最优操作是：
(1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14


提示：
1 <= n <= 7
nums.length == 2 * n
1 <= nums[i] <= 106
*/
