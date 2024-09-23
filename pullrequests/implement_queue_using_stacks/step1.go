//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
時間：15分

スタックを２つ使ってキューを実装する方法を知っていたので解法自体はすぐに思いついた。
LeetCodeの問題としてはPopやPeekは必ず有効な時に呼ばれるとのことなので、エラー処理をする必要はなく、プロトタイプ宣言自体もerrorを返せるようになっていないが、個人的にエラー処理をしないと気持ち悪かったので、変なエラー処理っぽいことをした中途半端なコードになってしまった。

末尾についているStep1は提出用につけただけなので無視してください。
*/
type MyQueueStep1 struct {
	pushStack []int
	popStack  []int
}

func ConstructorStep1() MyQueueStep1 {
	var q MyQueueStep1
	return q
}

func (this *MyQueueStep1) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueueStep1) Pop() int {
	this.Peek()
	if len(this.popStack) > 0 {
		x := this.popStack[len(this.popStack)-1]
		this.popStack = this.popStack[:len(this.popStack)-1]
		return x
	}
	return 0
}

func (this *MyQueueStep1) Peek() int {
	if len(this.popStack) <= 0 {
		for len(this.pushStack) > 0 {
			x := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
			this.popStack = append(this.popStack, x)
		}
	}
	if len(this.popStack) > 0 {
		return this.popStack[len(this.popStack)-1]
	}
	return 0
}

func (this *MyQueueStep1) Empty() bool {
	return len(this.pushStack) <= 0 && len(this.popStack) <= 0
}
