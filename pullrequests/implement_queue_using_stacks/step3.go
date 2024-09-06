//lint:file-ignore U1000 Ignore all unused code
package template

import "fmt"

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
実際にLeetCodeの制約を無視できる場合のコードも書いてみた（本来であれば`container/list`を使えば良いという話だが、そこは一応LeetCodeで提示されている形から極力変えずに書いた）。
*/
type Queue struct {
	pushStack []int
	popStack  []int
}

func (q *Queue) Push(n int) {
	q.pushStack = append(q.pushStack, n)
}

func (q *Queue) Pop() (int, error) {
	n, err := q.Peek()
	if err != nil {
		return 0, fmt.Errorf("queue is empty, cannot pop")
	}
	q.popStack = q.popStack[:len(q.popStack)-1]
	return n, nil
}

func (q *Queue) Peek() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("queue is empty, cannot peek")
	}
	if len(q.popStack) == 0 {
		for len(q.pushStack) > 0 {
			q.popStack = append(q.popStack, q.pushStack[len(q.pushStack)-1])
			q.pushStack = q.pushStack[:len(q.pushStack)-1]
		}
	}
	return q.popStack[len(q.popStack)-1], nil
}

func (q *Queue) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}
