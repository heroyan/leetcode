package problems

import "fmt"
//https://leetcode-cn.com/problems/climbing-stairs/submissions/

var mp  map[int]int = map[int]int{}

func climbStairs(n int) int {
	v, ok := mp[n]
	if ok {
		return v
	}

	if n < 3 {
		mp[n] = n
		return n
	}

	v = climbStairs(n-1) + climbStairs(n-2)
	mp[n] = v

	return v
}

func P70()  {
	fmt.Println(climbStairs(100))
	fmt.Println(climbStairs(200))
}
