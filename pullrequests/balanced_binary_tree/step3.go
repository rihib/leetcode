//lint:file-ignore U1000 Ignore all unused code
package template

import "math"

/*
Go styleguideを見ると、下記のように書かれているのでポインタではなく、直接構造体を渡すように直した。

> Do not pass pointers as function arguments just to save a few bytes. If a function reads its argument x only as *x throughout, then the argument shouldn’t be a pointer. Common instances of this include passing a pointer to a string (*string) or a pointer to an interface value (*io.Reader). In both cases, the value itself is a fixed size and can be passed directly.
>
> This advice does not apply to large structs, or even small structs that may increase in size. In particular, protocol buffer messages should generally be handled by pointer rather than by value. The pointer type satisfies the proto.Message interface (accepted by proto.Marshal, protocmp.Transform, etc.), and protocol buffer messages can be quite large and often grow larger over time.
*/
type treeBalance struct {
	isBalanced bool
	height     int
}

func isBalancedStep3(root *TreeNode) bool {
	return checkBalanceStep3(root).isBalanced
}

func checkBalanceStep3(root *TreeNode) treeBalance {
	if root == nil {
		return treeBalance{true, 0}
	}
	left, right := checkBalanceStep3(root.Left), checkBalanceStep3(root.Right)
	isBalanced := left.isBalanced && right.isBalanced &&
		math.Abs(float64(left.height-right.height)) <= 1.0
	height := max(left.height, right.height) + 1
	return treeBalance{isBalanced, height}
}
