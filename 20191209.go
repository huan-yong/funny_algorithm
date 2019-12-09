// author: huan-yong
// date: 2019.12.09 21:50:00
// 功能: 并发打印奇偶数，使得最终输出有序


package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	// 起一个go-routine打印奇数
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			if i > n {
				break
			}
			fmt.Println(i)
			// 通过两个chan完成唤醒&等待操作
			c2 <- struct{}{}
			<-c1
		}
	}(100)

	// 起一个go-routine打印偶数
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			if i > n {
				break
			}
			<-c2
			fmt.Println(i)
			c1 <- struct{}{}
		}
	}(100)

	// 主协程通过WaitGroup等待两个子协程结束
	wg.Wait()
}
