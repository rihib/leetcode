//lint:file-ignore U1000 Ignore all unused code
package diameterofbinarytree

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
Step1では厳密にリーフノードを深さ0として扱っていたが、深さ１として扱うことでコードをより単純にした。また命名も改善した。

`ml`という変数名については、`maxLength`とすると構造体名と被ってしまうのと、スコープが２行なので、`ml`でも許容できると思い、この変数名にした。

https://google.github.io/styleguide/go/decisions#getters にはgetterの関数名の先頭にGetというプレフィックスをつけるべきではないと書かれているのでそれに従っている。

また、maxDiameterで値渡しをしているのは、下記に書かれている通り、数バイトを節約するためだけに参照渡しをするなと書かれているのでそれに従っている。
https://google.github.io/styleguide/go/decisions#pass-values

l, rについては、スコープが３行なので、逆に変数名を長くすると可読性を下げると思い、１文字変数を使用している。
*/
func diameterOfBinaryTreeStep2(root *TreeNode) int {
	ml := maxDiameterStep2(root)
	return ml.diameter
}

type maxLength struct {
	diameter int
	depth    int
}

func maxDiameterStep2(root *TreeNode) maxLength {
	if root == nil {
		return maxLength{0, 0}
	}
	l, r := maxDiameterStep2(root.Left), maxDiameterStep2(root.Right)
	diameter := max(l.diameter, r.diameter, l.depth+r.depth)
	depth := max(l.depth, r.depth) + 1
	return maxLength{diameter, depth}
}
