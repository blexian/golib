package tree

import "fmt"

type node struct {
	Value int
	left  *node
	right *node
}

// [ 0 10 11 21 -1 -1 -1] 代表以下二叉树结构，负数表示节点为nil
/*
		   0
		  / \
		10   11
	   /
	  21
*/
// [ 0 10 11 20 21 -1 23 30 -1 -1 33 -1 -1 36 37 ]
/*
                   0
				  / \
				10   11
			   /  \    \
			  20  21    23
             /     \   /  \
            30     33 36  37
*/
func constructor(arr []int) *node {
	// arr[i]的左子为arr[2*i+1], 右子为arr[2*i+2]
	// 数组的长度一定为 2^n-1，n为层数
	return consNode(arr, 0)
}

func consNode(arr []int, index int) *node {
	l := len(arr)
	if index >= l {
		return nil
	}
	if arr[index] < 0 {
		return nil
	} else {
		return &node{
			Value: arr[index],
			left:  consNode(arr, 2*index+1),
			right: consNode(arr, 2*index+2),
		}
	}
}

// breadthFirstSearch 广度优先算法
// 思路：
// 1、访问根，将根节点的所有子节点放入队列queue中
// 2、依次访问queue中节点，并将节点的孩子放入queue中
// 3、重复第2步操作，知道队列为空
func breadthFirstSearch(n *node) {
	fmt.Printf("[ ")
	queue := make([]*node, 0)
	// 入队列
	queue = append(queue, n)
	for len(queue) > 0 && queue[0] != nil {
		// 出队列
		nd := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", nd.Value)
		if nd.left != nil {
			queue = append(queue, nd.left)
		}
		if nd.right != nil {
			queue = append(queue, nd.right)
		}
	}
	fmt.Printf("]\n")
}

// DepthFirstPreSearch 深度优先算法 -- 前序
func DepthFirstPreSearch(n *node) {
	fmt.Printf("[ ")
	if n != nil {
		depthFirstPreSearch(n)
	}
	fmt.Printf("]\n")
}
func depthFirstPreSearch(n *node) {
	if n.left != nil {
		depthFirstPreSearch(n.left)
	}
	fmt.Printf("%d ", n.Value)
	if n.right != nil {
		depthFirstPreSearch(n.right)
	}
}

// DepthFirstMidSearch 深度优先算法 -- 中序
func DepthFirstMidSearch(n *node) {
	fmt.Printf("[ ")
	if n != nil {
		depthFirstMidSearch(n)
	}
	fmt.Printf("]\n")
}

func depthFirstMidSearch(n *node) {
	fmt.Printf("%d ", n.Value)
	if n.left != nil {
		depthFirstMidSearch(n.left)
	}
	if n.right != nil {
		depthFirstMidSearch(n.right)
	}
}

// DepthFirstPostSearch 深度优先算法 -- 后序
func DepthFirstPostSearch(n *node) {
	fmt.Printf("[ ")
	if n != nil {
		depthFirstPostSearch(n)
	}
	fmt.Printf("]\n")
}

func depthFirstPostSearch(n *node) {
	if n.right != nil {
		depthFirstPostSearch(n.right)
	}
	fmt.Printf("%d ", n.Value)
	if n.left != nil {
		depthFirstPostSearch(n.left)
	}
}
