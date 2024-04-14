//lint:file-ignore U1000 Ignore all unused code
package main

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		tmp := root.Left
		root.Left, root.Right = invertTree(root.Right), invertTree(tmp)
	}
	return root
}
