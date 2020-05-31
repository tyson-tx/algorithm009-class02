
// 242. 有效的字母异位词
// link: https://leetcode-cn.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 创建两个仿 bitmap 的数组
	bitmapS := [26]int{}
	bitmapT := [26]int{}
	for i := 0; i < len(s); i++ {
		// 使用 ascii 码计算对应仿 bitmap 数组的下标
		index := s[i]-'a'
		bitmapS[index]++
	}
	for i := 0; i < len(t); i++ {
		index := t[i]-'a'
		bitmapT[index]++
	}
	// 最后比较两个数组
	return bitmapT == bitmapS
}