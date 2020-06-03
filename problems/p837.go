package problems

// 新21点 https://leetcode-cn.com/problems/new-21-game/

import (
	"fmt"
)

var total = 0
var count = 0
var count2 = 0

func new21Game(N int, K int, W int) float64 {
	// 从后往前推，假如最后一步是1, 前面所有步数的和必须是k-1
	// 假如最后一步是2，前面所有步数的和可以是k-1和k-2
	// 假如最后一步是n，前面所有步数的和可以是k-n, k-n+1, k-n+2..满足k-n+j>0 且
	numkw(N, K, W)

	return float64(count2) / float64(count)
}

func numkw(N, K, W int)  {
	if K <= 0 {
		count++
		if total <= N {
			count2++
		}
		return
	}

	for i := 1; i <= W; i++ {
		total += i
		numkw(N, K-total, W)
		total -= i
	}
}

func P837()  {
	fmt.Println(new21Game(6, 1, 10))
	fmt.Println(new21Game(21, 17, 10))
}