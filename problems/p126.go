package problems

import (
	"fmt"
)

// https://leetcode-cn.com/problems/word-ladder-ii/

func findLadders(beginWord, endWord string, wordList []string) [][]string {
	if beginWord == endWord {
		return [][]string{}
	}

	ret := [][]string{}
	nextList := nextWord(beginWord, wordList)
	for _, w := range nextList {
		nextarr := findLadders(w, endWord, wordList)
		fmt.Println(nextarr)
		for _, iarr := range nextarr {
			ret = append(ret, []string{beginWord}, iarr)
		}
	}

	return ret
}

func nextWord(word string, wordList []string) (ret []string) {
	for _, w := range wordList {
		dist, _ := distance(word, w)
		if dist == 1 {
			ret = append(ret, w)
		}
	}
	return
}

func distance(word1, word2 string) (dist int, index int) {
	index = -1
	for idx, _ := range word1 {
		if word2[idx] != word1[idx] {
			dist++
			index = idx
		}
	}
	return
}

func P126()  {
	findLadders("hit", "dog", []string{"hot","dot","dog","lot","log","cog"})
}
