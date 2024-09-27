//lint:file-ignore U1000 Ignore all unused code
package diameterofbinarytree

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
時間：22分

解法自体はすぐに思いついたが、実装に時間がかかってしまった。特に命名が難しくて悩んでしまった。

解法としては、その時の最大のdiameterと、左の深さ+右の深さのうち大きい方で最大のdiameterを更新していけばいい。
*/
func diameterOfBinaryTreeStep1(root *TreeNode) int {
	info := maxDiameterStep1(root)
	return info.maxLength
}

type pathInfo struct {
	depth     int
	maxLength int
}

func maxDiameterStep1(root *TreeNode) pathInfo {
	if root == nil || root.Left == nil && root.Right == nil {
		return pathInfo{0, 0}
	}
	leftInfo := maxDiameterStep1(root.Left)
	rightInfo := maxDiameterStep1(root.Right)

	var leftDepth, rightDepth int
	if root.Left == nil {
		leftDepth = leftInfo.depth
	} else {
		leftDepth = leftInfo.depth + 1
	}
	if root.Right == nil {
		rightDepth = rightInfo.depth
	} else {
		rightDepth = rightInfo.depth + 1
	}
	depth := max(leftDepth, rightDepth)
	maxLength := max(leftDepth+rightDepth, leftInfo.maxLength, rightInfo.maxLength)
	return pathInfo{depth, maxLength}
}
