/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    
	var dfs func(node, lower, upper *TreeNode) bool

	dfs = func(node, lower, upper *TreeNode) bool{

		if node == nil{
			return true
		}


		if lower != nil && node.Val <= lower.Val{
			return false
		}

		if upper != nil && node.Val>= upper.Val{
			return false 
		}


		return dfs(node.Left, lower, node) && dfs(node.Right, node, upper)
	}

	return dfs(root, nil, nil)
}
