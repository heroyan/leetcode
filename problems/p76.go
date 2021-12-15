package problems

import "fmt"

func minWindow(s string, t string) string {
	needs := map[uint8]int{}
	window := map[uint8]int{}
	slen := len(s)
	tlen := len(t)
	for i := 0; i < tlen; i++ {
		needs[t[i]]++
	}
	var left int
	var right int
	var valid int
	var start = -1
	var maxlen = 999999999999
	for right < slen {
		if needs[s[right]] > 0 {
			window[s[right]]++
			// 出现相等的时候表示一个字母的个数已经满足了，后面可能还是出现，于是有可能  window[s[right]] > needs[s[right]]
			if window[s[right]] == needs[s[right]] {
				valid++
			}
		}
		right++
		for valid == len(needs) {
			if right - left < maxlen {
				start = left
				maxlen = right - left
			}
			if needs[s[left]] > 0 {
				// 字符数减到刚好等于needs数的时候，说明再减这个就要不符合条件了，前面可能出现比needs数更多的同字符
				if window[s[left]] == needs[s[left]] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}

	if start < 0 {
		return ""
	}

	return s[start:start+maxlen]
}

func P76()  {
	fmt.Println(minWindow("jjjjjjkksfasd", "abcd"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("aa", "a"))
}
