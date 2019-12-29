// author: huan-yong
// date: 2019.12.29 09:53:00
// 功能: 将正整数数字转换为16进制格式输出

package main

import (
	"fmt"
)

func f(n int) string {
	res := ""
	if n <= 0 {
		return res
	}

	for n > 0 {
		c := n % 16
		n = n / 16
		d := fmt.Sprintf("%d", c)
		if c == 10 {
			d = "A"
		} else if c == 11 {
			d = "B"
		} else if c == 12 {
			d = "C"
		} else if c == 13 {
			d = "D"
		} else if c == 14 {
			d = "E"
		} else if c == 15 {
			d = "F"
		}

		res = d + res
	}

	return res
}

func main() {
	fmt.Println(f(31))
}
