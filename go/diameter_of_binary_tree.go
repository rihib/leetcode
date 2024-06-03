//lint:file-ignore U1000 Ignore all unused code
package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	getmaxLen(root, &maxLength)
	return maxLength
}

func getmaxLen(t *TreeNode, maxLength *int) int {
	if t == nil {
		return 0
	}

	left, right := getmaxLen(t.Left, maxLength), getmaxLen(t.Right, maxLength)
	*maxLength = max(*maxLength, left+right)

	return max(left, right) + 1
}
