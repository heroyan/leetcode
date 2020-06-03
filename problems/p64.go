package problems

import "fmt"
//https://leetcode-cn.com/problems/qiu-12n-lcof/

func sumNums(n int) int {
    ans := 0
    var sum func(int) bool
    sum = func(nn int) bool {
        ans += nn
        return nn > 0 && sum(nn-1)
    }

    sum(n)

    return ans
}

func sumNums2(n int) int {
    ans := 0
    a := n
    b := n+1
    add := func() bool {
    	ans += a
    	return true
	}

	_ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    _ = (b&1 == 1) && add()
    a <<= 1
    b >>= 1

    // 公式：n*(n+1)/2

    return ans >> 1
}

func P64()  {
	fmt.Println(sumNums2(100))
	fmt.Println(sumNums(100))
}

