//lint:file-ignore U1000 Ignore all unused code
package blind75

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		tmp := root.Left
		root.Left = invertTree(root.Right)
		root.Right = invertTree(tmp)
	}
	return root
}
