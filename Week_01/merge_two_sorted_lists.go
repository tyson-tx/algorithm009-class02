// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
type ListNode struct {
	Val string
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fastNode := &ListNode{}
	result := fastNode // 为了保存头节点
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			fastNode.Next = l1
			l1 = l1.Next
		} else {
			fastNode.Next = l2
			l2 = l2.Next
		}
		fastNode = fastNode.Next
	}
	if l1 != nil {
		fastNode.Next = l1
	}
	if l2 != nil {
		fastNode.Next = l2
	}
	return result.Next
}