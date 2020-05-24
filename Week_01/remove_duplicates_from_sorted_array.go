// 26. 删除排序数组中的重复项
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates(nums []int) int {
	m := len(nums)
	if m < 2 {
		return m
	}
	start, end := 1, 1
	for end < m {
		if nums[end] != nums[end-1] {
			nums[start] = nums[end]
			start++
		}
		end++
	}
	return start
}