// 127. 单词接龙
// link：https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
    dict := make(map[string]bool)
    for _, word := range wordList {
        dict[word] = true
    }
    if _, ok := dict[endWord]; !ok {
        return 0
    }
    q1 := make(map[string]bool)
    q2 := make(map[string]bool)
    q1[beginWord] = true
    q2[endWord] = true
    l := len(beginWord)
    steps := 0
    for len(q1) > 0 && len(q2) > 0 {
        steps++
        if len(q1) > len(q2) {
            q1, q2 = q2, q1
        }
        q := make(map[string]bool)
        for k := range q1 {
            chs := []rune(k)
            for i := 0; i < l; i++ {
                ch := chs[i]
                for c := 'a'; c <= 'z'; c++ {
                    chs[i] = c
                    t := string(chs)
                    if _, ok := q2[t]; ok {
                        return steps+1
                    }
                    if _, ok := dict[t]; !ok {
                        continue
                    }
                    delete(dict, t)
                    q[t] = true
                }
                chs[i] = ch
            }
        }
        q1 = q
    }
    return 0
}