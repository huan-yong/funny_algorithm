// author: huan-yong
// date: 2019.12.15 20:06:00
// 功能: 蓄水池抽样算法

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reservoirSampling(a []int64, k int) []int64 {
	if a == nil || len(a) == 0 || k <= 0 {
		return []int64{}
	}

	n := k
	if len(a) < n {
		n = len(a)
	}

	ret := make([]int64, n)
	rand.Seed(time.Now().UnixNano())
	for i, v := range a {
		// 前k项默认构成一个蓄水池
		if i < k {
			ret[i] = v
		} else {
			rd := rand.Intn(i)
			// 以k/i的概率替换第rd项
			if rd < k {
				ret[rd] = v
			}
		}
	}

	return ret
}

func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := make([]int, 10 + 1)
	for i := 0; i < 10000; i++ {
		ret := reservoirSampling(a, 3)
		for _, v := range ret {
			// 统计每个数被抽中的频次
			b[v]++
		}
	}

	for i, v := range b {
		if v > 0 {
			fmt.Printf("%d:%d\n", i, v)
		}
	}
}
