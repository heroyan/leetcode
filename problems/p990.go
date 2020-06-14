package problems

import (
	"fmt"
)

// https://leetcode-cn.com/problems/satisfiability-of-equality-equations/

func equationsPossible(equations []string) bool {
	mp := map[byte]int{}
	gmp := map[int][]byte{}
	groupid := 1
	// 先把相等的归类
	for _, eq := range equations {
		v1, ok1 := mp[eq[0]]
		v2, ok2 := mp[eq[3]]
		if eq[1] == '=' {
			if ok1 && ok2 {
				// 两者合并分组
				mp[eq[3]] =  v1
				for _, b := range gmp[v2]  {
					mp[b] = v1
				}
				gmp[v1] = append(gmp[v1], gmp[v2]...)
			} else if ok1 && !ok2 {
				mp[eq[3]] =  v1
				gmp[v1] = append(gmp[v1], eq[3])
			} else if !ok1 && ok2 {
				mp[eq[0]] =  v2
				gmp[v2] = append(gmp[v2], eq[0])
			} else {
				mp[eq[0]] = groupid
				mp[eq[3]] = groupid
				groupid++
			}
		}
	}

	for _, eq := range equations {
		//v1, ok1 := mp[eq[0]]
		//v2, ok2 := mp[eq[3]]
		if eq[1] == '!' {

		}
	}

	return true
}

func P990()  {
	fmt.Println(equationsPossible([]string{"a==b","b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a","a==b"}))
	fmt.Println(equationsPossible([]string{"a==b","b==c","a==c"}))
	fmt.Println(equationsPossible([]string{"a==b","b!=c","c==a"}))
	fmt.Println(equationsPossible([]string{"c==c","b==d","x!=z"}))
}
