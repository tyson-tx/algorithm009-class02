// 127. 单词接龙
// link：https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
    // 判断边界
    var flag bool
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord {
            flag = true
            break
        }
    }
    if !flag {
        return 0
    }

    var res int
    lqueue, rqueue := []string{beginWord}, []string{endWord}
    visited := make(map[string]bool)
    for len(lqueue) != 0 {
        res++
        tmpqueue := []string{}
        size := len(lqueue)
        for size > 0 {
            str := lqueue[0]
            lqueue = lqueue[1:]
            // 判断在另个队列是否已经存在，如果存在，则返回
            for i := 0; i < len(rqueue); i++ {
                if compare(str, rqueue[i]) {
                    return res+1
                }
            }
            // 如果不存在，则从未访问过的单词中选择加入队列中
            for i := 0; i < len(wordList); i++ {
                if _, ok := visited[wordList[i]]; !ok && compare(str, wordList[i]) {
                    tmpqueue = append(tmpqueue, wordList[i])
                    visited[wordList[i]] = true
                }
            }
            size--
        }
        lqueue = tmpqueue
        if len(lqueue) >= len(rqueue) {
            lqueue, rqueue = rqueue, lqueue
        }
    }
    return 0
}

func compare(str1, str2 string) bool {
    var count int
    for i := 0; i < len(str1); i++ {
        if str1[i] != str2[i] {
            count++
        }
    }
    return count==1
}