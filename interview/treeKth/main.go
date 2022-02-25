package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 0, 2}
	tree := createBinaryTree(0, arr)
	fmt.Println(kthLargest(tree, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTraverse(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	inorderTraverse(root.Left, ret)
	*ret = append(*ret, root.Val)
	inorderTraverse(root.Right, ret)
}

func kthLargest(root *TreeNode, k int) int {
	ret := &[]int{}
	inorderTraverse(root, ret)
	// 说明没找到第k大节点
	if k < 1 && k > len(*ret) {
		return -1
	}

	return (*ret)[len(*ret)-k]
}

func createBinaryTree(i int, nums []int) *TreeNode {
	tree := &TreeNode{nums[i], nil, nil}
	// 左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = createBinaryTree(2*i+1, nums)
	}
	// 右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = createBinaryTree(2*i+2, nums)
	}
	return tree
}
