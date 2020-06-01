package problems

import "fmt"

func divide(dividend int, divisor int) int {
    if dividend == 0 {
		return 0
	}

    if divisor == 1 {
        return dividend
    }

	max := 1 << 31 - 1
	min := -1 << 31

	if divisor == min {
		if dividend == min {
			return 1
		} else {
			return 0
		}
	}

	if divisor == -1 {
		if dividend == min {
			return max
		} else {
			return -dividend
		}
	}

	unsigned := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		unsigned = false
	}

	// 变成正数
	if dividend < 0 {
        dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

    if dividend < divisor {
        return 0
    }

	count := div(dividend, divisor)
	if !unsigned {
		return -count
	}

	return count
}

func div(a, b int) int {
    if a < b {
        return 0
    }
    count := 1
    t := b
    for t + t <= a {
        count = count + count
        t = t + t
    }

    return count + div(a-t, b)
}

func Div()  {
	fmt.Println(divide(100, 12))
	fmt.Println(divide(12, 12))
	fmt.Println(divide(-112, 12))
}
