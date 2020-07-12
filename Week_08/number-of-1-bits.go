// 191. 位 1 的个数
// link：https://leetcode-cn.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
    count := 0
    for num > 0 {
        if num % 2 == 1 {
            count++
        }
        num /= 2
    }
    return count
}