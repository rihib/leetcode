//lint:file-ignore U1000 Ignore all unused code
package template

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
LeetCodeに通すことだけを考えてちゃんと綺麗に書いてみた。
*/
type MyQueueStep2 struct {
	pushStack []int
	popStack  []int
}

func ConstructorStep2() MyQueueStep2 {
	return MyQueueStep2{}
}

func (q *MyQueueStep2) Push(n int) {
	q.pushStack = append(q.pushStack, n)
}

func (q *MyQueueStep2) Pop() int {
	n := q.Peek()
	q.popStack = q.popStack[:len(q.popStack)-1]
	return n
}

func (q *MyQueueStep2) Peek() int {
	if len(q.popStack) == 0 {
		for len(q.pushStack) > 0 {
			q.popStack = append(q.popStack, q.pushStack[len(q.pushStack)-1])
			q.pushStack = q.pushStack[:len(q.pushStack)-1]
		}
	}
	return q.popStack[len(q.popStack)-1]
}

func (q *MyQueueStep2) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}
