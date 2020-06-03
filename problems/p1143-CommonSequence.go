package problems

import "fmt"
// https://leetcode-cn.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
	return 0
}

func longestCommonStr(text1, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)

	if len1 == 0 || len2 == 0 {
		return 0
	}

	s := [][]int{}
	for i := 0; i < len1; i++ {
		s = append(s, make([]int, len2))
	}



	fmt.Println(s)

	return 0
}

func P1143()  {
	fmt.Println(longestCommonStr("fadfaf", "fasfasd"))
}