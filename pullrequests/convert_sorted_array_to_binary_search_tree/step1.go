//lint:file-ignore U1000 Ignore all unused code
package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：10分
最初に解法を思いつくのに時間がかかったが、BSTはノードの値がleft < root < rightということを思い出して解くことができた。

Goではスライスはポインタをずらすだけなのでメモリコピーは発生しないという認識。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	return &TreeNode{
		Val:   nums[midIndex],
		Left:  sortedArrayToBST(nums[:midIndex]),
		Right: sortedArrayToBST(nums[midIndex+1:]),
	}
}
