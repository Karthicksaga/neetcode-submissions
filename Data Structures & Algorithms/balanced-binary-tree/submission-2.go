/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func isBalanced(root *TreeNode) bool {
    
	var dfs func(root *TreeNode) int 

	dfs = func(root *TreeNode) int{

		if root == nil {
			return 0 
		}

		left := dfs(root.Left)
		if left == -1 {
			return -1
		}

		right := dfs(root.Right)
		if right == -1{
			return -1 
		}

	  	if abs(left - right) > 1 {
        	return -1
    	}


		return 1 + max(left, right)
	}

	return dfs(root) != -1
}
