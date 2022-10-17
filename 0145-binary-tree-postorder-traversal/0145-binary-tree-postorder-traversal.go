/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postOrder(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }
    
    postOrder(root.Left, arr)
    postOrder(root.Right, arr)
    *arr = append(*arr, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
    arr := []int{}
    postOrder(root, &arr)
    return arr
}

