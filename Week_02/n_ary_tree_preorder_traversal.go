// 589. N叉树的前序遍历
// link: https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val int
	Children []*Node
}
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val) // 前序输出
			if len(root.Children) == 0 {
				break
			}
			for i := len(root.Children)-1; i > 0; i-- {
				stack = append(stack, root.Children[i]) // 入栈
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

