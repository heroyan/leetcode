package problems

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	needs := map[uint8]int{}
	window := map[uint8]int{}
	slen := len(s)
	tlen := len(p)
	for i := 0; i < tlen; i++ {
		needs[p[i]]++
	}
	var left int
	var right int
	var valid int
	ret := []int{}
	for right < slen {
		if needs[s[right]] > 0 {
			window[s[right]]++
			// 出现相等的时候表示一个字母的个数已经满足了，后面可能还是出现，于是有可能  window[s[right]] > needs[s[right]]
			if window[s[right]] == needs[s[right]] {
				valid++
			}
		}
		right++
		for right - left >= tlen {
			if valid == len(needs) {
				ret = append(ret, left)
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

	return ret
}