// 493. 翻转对
// link：https://leetcode-cn.com/problems/reverse-pairs/

// 归并排序
func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(arr []int, l int, r int) int {
    if l >= r {
        return 0
    }
    count, mid := 0, (l+r)>>1
    count = mergeSort(arr, l, mid) + mergeSort(arr, mid+1, r)
    temp, i, k := make([]int, r-l+1), l, 0
    for j, index := mid+1, l; j <= r; j++ {
        for ; i <= mid && arr[i] < arr[j]; i++ {
            temp[k], k = arr[i], k+1
        }
        temp[k], k = arr[j], k+1
        for index <= mid && (arr[index]+1)>>1 <= arr[j] {
            index++
        }
        count += mid-index+1
    }
    copy(arr[l+k:], arr[i:mid+1])
    copy(arr[l:], temp[:k])
    return count
}