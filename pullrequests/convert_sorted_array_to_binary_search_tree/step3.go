//lint:file-ignore U1000 Ignore all unused code
package convertsortedarraytobinarysearchtree

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
その他にも最初にdummyノードで完全二分木を構築してから、In-orer traversalで値をセットしていくという方法で解いている人もいたので、同様に実装してみた。
*/
func sortedArrayToBSTInorder(nums []int) *TreeNode {
	root := dummyCompleteBinaryTree(0, len(nums))
	setValueInorder(root, nums)
	return root
}

func dummyCompleteBinaryTree(index, length int) *TreeNode {
	if index >= length {
		return nil
	}
	node := &TreeNode{}
	node.Left = dummyCompleteBinaryTree(index*2+1, length)
	node.Right = dummyCompleteBinaryTree(index*2+2, length)
	return node
}

func setValueInorder(root *TreeNode, nums []int) []int {
	if root == nil || len(nums) == 0 {
		return nums
	}
	nums = setValueInorder(root.Left, nums)
	root.Val = nums[0]
	nums = setValueInorder(root.Right, nums[1:])
	return nums
}
