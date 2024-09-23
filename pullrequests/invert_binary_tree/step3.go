//lint:file-ignore U1000 Ignore all unused code
package invertbinarytree

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTreeRecursive(root.Right), invertTreeRecursive(root.Left)
	return root
}
