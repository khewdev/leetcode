/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, left, right int) bool {
    if root == nil {
        return true
    }
    
    if root.Val <= left || root.Val >= right {
        return false
    }
    
    return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}