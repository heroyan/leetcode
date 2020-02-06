//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
    slen := len(s)

    var i int = 0
    var j int = 0
    var max int = 0
    var cmap = make(map[byte] bool)

    for {
        if i >= slen || j >= slen {
            break
        }
        _, ok := cmap[s[j]]
        if !ok {
            tmp := j - i + 1
            if tmp > max {
                max = tmp
            }
            cmap[s[j]] = true
            j++
        } else {
            delete(cmap, s[i])
            i++
        }
    }

    return max
}
