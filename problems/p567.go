package problems

//https://leetcode-cn.com/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	needs := map[uint8]int{}
	window := map[uint8]int{}
	slen := len(s2)
	tlen := len(s1)
	for i := 0; i < tlen; i++ {
		needs[s1[i]]++
	}
	var left int
	var right int
	var valid int
	for right < slen {
		if needs[s2[right]] > 0 {
			window[s2[right]]++
			// 出现相等的时候表示一个字母的个数已经满足了，后面可能还是出现，于是有可能  window[s[right]] > needs[s[right]]
			if window[s2[right]] == needs[s2[right]] {
				valid++
			}
		}
		right++
		for right - left >= tlen {
			if valid == len(needs) {
				return true
			}
			if needs[s2[left]] > 0 {
				// 字符数减到刚好等于needs数的时候，说明再减这个就要不符合条件了，前面可能出现比needs数更多的同字符
				if window[s2[left]] == needs[s2[left]] {
					valid--
				}
				window[s2[left]]--
			}
			left++
		}
	}

	return false
}