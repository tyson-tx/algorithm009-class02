// 151. 翻转字符串里的单词
// link：https://leetcode-cn.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
    list := strings.Split(s, " ")
    var res []string
    for i := len(list)-1; i >= 0; i-- {
        if len(list[i]) > 0 {
            res = append(res, list[i])
        }
    }
    s = strings.Join(res, " ")
    return s
}