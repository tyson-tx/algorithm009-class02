package main

import (
	"sort"
)

// 49. 字母异位词分组
// link: https://leetcode-cn.com/problems/group-anagrams/

type bytes []byte
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := make(map[string]int)
	for _, str := range strs {
		kByte := bytes(str)
		sort.Sort(kByte)
		k := string(kByte)
		if idx, ok := m[k]; !ok {
			m[k] = len(ret)
			ret = append(ret, []string{str})
		} else {
			ret[idx] = append(ret[idx], str)
		}
	}
	return ret
}