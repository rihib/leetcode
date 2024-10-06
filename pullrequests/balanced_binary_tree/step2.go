//lint:file-ignore U1000 Ignore all unused code
package template

import "math"

/*
読みやすくなるようにリファクタした。
Go的にはleftBalance.isBalancedのように同じ名前が繰り返されるのを嫌う（Go styleguideにも書いてある）ので、繰り返さないように直したりした。

正直、TreeBalanceやcheckBalanceという関数名はしっくりきていないけど、より良い名前が思いつかなかった。
*/
type TreeBalance struct {
	isBalanced bool
	height     int
}

func isBalancedStep2(root *TreeNode) bool {
	return checkBalanceStep2(root).isBalanced
}

func checkBalanceStep2(root *TreeNode) *TreeBalance {
	if root == nil {
		return &TreeBalance{true, 0}
	}
	left, right := checkBalanceStep2(root.Left), checkBalanceStep2(root.Right)
	isBalanced := left.isBalanced && right.isBalanced &&
		math.Abs(float64(left.height-right.height)) <= 1.0
	height := max(left.height, right.height) + 1
	return &TreeBalance{isBalanced, height}
}
