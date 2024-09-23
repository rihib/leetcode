//lint:file-ignore U1000 Ignore all unused code
package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
時間：3分30秒

再帰的に左右のノードを入れ替えていけば良いというのにはすぐに気づき、すんなり実装することができた。
*/
func invertTreeStep1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTreeStep1(root.Left)
	invertTreeStep1(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
