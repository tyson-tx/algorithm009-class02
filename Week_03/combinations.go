// 77. 组合
// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// link：https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
    if n == 0 || k == 0 {
        return [][]int{}
    }
    var res [][]int
    var elem []int
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i+1
    }
    combineRec(nums, 0, k, elem, &res)
    return res
}
func combineRec(nums []int, start, k int, elem []int, result *[][]int) {
    if len(elem) == k {
        tmp := make([]int, k)
        copy(tmp, elem)
        *result = append(*result, tmp)
    }
    for i := start; i < len(nums); i++ {
        elem = append(elem, nums[i])
        combineRec(nums, i+1, k, elem, result)
        elem = elem[:len(elem)-1]
    }
}


