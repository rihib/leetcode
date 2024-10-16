//lint:file-ignore U1000 Ignore all unused code
package template

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
時間：12分

左右の木の高さを比べてバランスしているかを確かめるのと、そもそもその部分木自体がバランスしているのかも確かめる必要があるので、別途再帰用の補助関数を作ることにした。
*/
type Balance struct {
	isBalanced bool
	height     int
}

func isBalancedStep1(root *TreeNode) bool {
	return isHeightBalanced(root).isBalanced
}

func isHeightBalanced(root *TreeNode) *Balance {
	if root == nil {
		return &Balance{true, 0}
	}
	leftBalance := isHeightBalanced(root.Left)
	rightBalance := isHeightBalanced(root.Right)
	isBalanced := leftBalance.isBalanced && rightBalance.isBalanced && (math.Abs(float64(leftBalance.height-rightBalance.height)) <= 1.0)
	return &Balance{isBalanced, max(leftBalance.height, rightBalance.height) + 1}
}
