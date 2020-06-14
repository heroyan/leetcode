package problems

import "fmt"
//https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	strlen := len(strs)

	if strlen == 0 {
		return ""
	}

	if strlen == 1 {
		return strs[0]
	}

	shortest := 1 << 31
	for _, v := range strs {
		lenv := len(v)
		if lenv < shortest {
			shortest = lenv
		}
	}

	result := []byte{}
	for i := 0; i < shortest; i++ {
		flag := false
		for j := 1; j < strlen; j++ {
			if strs[j][i] != strs[j-1][i] {
				flag = true
				break
			}
		}

		if flag {
			break
		}

		result = append(result, strs[0][i])
	}

	return string(result)
}

func P14()  {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"flower","blow","alight"}))
}
