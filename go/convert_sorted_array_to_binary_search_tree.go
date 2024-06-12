//lint:file-ignore U1000 Ignore all unused code
package main

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildBST(nums, left, mid-1)
	node.Right = buildBST(nums, mid+1, right)

	return node
}
