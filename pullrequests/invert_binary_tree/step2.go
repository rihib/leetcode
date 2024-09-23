//lint:file-ignore U1000 Ignore all unused code
package invertbinarytree

/*
Step1は冗長だったのでリファクタした。
これぐらい短いと早期リターンしなくても良いかなと思い下記のようにしました。
またイテレーティブに解く方法も実装しました。
*/
func invertTreeStep2(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTreeStep2(root.Right), invertTreeStep2(root.Left)
	}
	return root
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
