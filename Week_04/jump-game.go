// 55. 跳跃游戏
// link：https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
    length := len(nums) - 1
    for i := length - 1; i >= 0; i-- {
        if nums[i]+i >= length {
            length = i
        }
    }
    return length <= 0
}