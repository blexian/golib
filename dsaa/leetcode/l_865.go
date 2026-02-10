package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	dp := make(map[*TreeNode]*TreeNode, 0)
	// dp[i] 表示i节点的父亲节点
	depNodes := make(map[int][]*TreeNode, 0)
	// 记录节点的深度
	var dfs func(*TreeNode, *TreeNode, int)
	var maxDepth int
	dfs = func(node *TreeNode, fa *TreeNode, depth int) {
		if node == nil {
			return
		}
		dp[node] = fa
		if depNodes[depth] == nil {
			depNodes[depth] = make([]*TreeNode, 0)
		}
		if depth > maxDepth {
			maxDepth = depth
		}
		depNodes[depth] = append(depNodes[depth], node)
		dfs(node.Left, node, depth+1)
		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 1)
	// 找到最深节点的所有父亲节点，然后求交集
	return root
}

// i节点是否包含j
// dp[i] = dp[i.left] || dp[i.right]
func containNode(i, j *TreeNode) bool {

	return false
}

func bfsFindDeepNodes(root *TreeNode) []*TreeNode {
	return nil
}
