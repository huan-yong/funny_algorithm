// author: huan-yong
// date: 2019.12.23 09:44:00
// 功能: n个苹果放到m个盘子里，允许有空盘，(5,1,1)和
// (1,5,1)看作是同一种放置方法

package main

import (
	"fmt"
)

func f(m, n int) int {
	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	for i := 0; i <= m; i++ {
		// i个苹果放到0个盘子里，方法数初始化为0
		dp[i][0] = 0
		// i个苹果放到1个盘子里，方法数初始化为1
		dp[i][1] = 1
	}

	for j := 0; j <= n; j++ {
		// 考虑到dp[i][j]=dp[i][j-1]+dp[i-j][j]中不存在空盘时
		// dp[i-j][j]是有值的，所以0个苹果放到j个盘子中，方法
		// 数初始化为1
		dp[0][j] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i < j {
				// 苹果数小于盘子数时，则至少有j-i个盘子为空盘
				// 所以i个苹果放到j个盘子里和放到i个盘子里方法
				// 数是一样的
				dp[i][j] = dp[i][i]
			} else {
				// 放置方法分为有一个空盘和无空盘两类
				dp[i][j] = dp[i][j - 1] + dp[i - j][j]
			}
		}
	}

	return dp[m][n]
}

func main() {
	m := 7
	n := 3

	fmt.Println(f(m, n))
}
