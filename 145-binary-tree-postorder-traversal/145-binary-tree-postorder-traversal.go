/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, arr)
	postorder(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
    arr := []int{}
	postorder(root, &arr)
    return arr
}