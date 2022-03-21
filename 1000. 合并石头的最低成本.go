package main

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + stones[i-1]
	}
	for l := k; l <= n;l++ { // 枚举区间长度
		for i := 1;i+l-1 <= n;i++ { // 枚举区间起点
			j := i + l - 1
			//i，j分别为区间左右端点
			dp[i][j] = math.MaxInt32
			for p := i; p < j;p += k - 1 { // 枚举分界点
				if dp[i][j]> dp[i][p]+dp[p+1][j]{
					dp[i][j] = dp[i][p]+dp[p+1][j]
				}
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += sum[j] - sum[i-1]
			}
		}
	}
	return dp[1][n]
}

func main() {
	stones :=[]int{3,2,4,1}
	K := 2
	println(mergeStones(stones,K))
}
