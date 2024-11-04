//lint:file-ignore U1000 Ignore all unused code
package main

type entry struct {
	node  *TreeNode
	depth int
}

/*
再帰
*/
// １フレームの大きさは、引数8 + 返り値8 + ローカル変数(8 + 8) + FP + 戻りアドレス = 48Bぐらい。
// 最大ノード数が10^4なのでスタックサイズは、48 * 10^4 = 480KBぐらいになると予想できる。
// Linuxのデフォルトのスタックの大きさが8MBなのでスタックオーバーフローの危険性は低い。
// goroutineのデフォルトのスタックの大きさは8KBだが、動的に拡張されるようになっているので通常、スタックサイズがGBレベルになっても問題なく再帰を回すことができる。
// https://github.com/rihib/leetcode/pull/15#issue-2459505166
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)
	return max(leftDepth, rightDepth) + 1
}

/*
DFS
*/
// pushする前にnilノードを弾く
func maxDepthIterativePreOrder1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	stack := []entry{{root, 1}}
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maximum = max(maximum, e.depth)
		if e.node.Right != nil {
			stack = append(stack, entry{e.node.Right, e.depth + 1})
		}
		if e.node.Left != nil {
			stack = append(stack, entry{e.node.Left, e.depth + 1})
		}
	}
	return maximum
}

// popした後にnilノードを弾く
func maxDepthIterativePreOrder2(root *TreeNode) int {
	maximum := 0
	stack := []entry{{root, 1}}
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if e.node == nil {
			continue
		}
		maximum = max(maximum, e.depth)
		stack = append(stack, entry{e.node.Right, e.depth + 1})
		stack = append(stack, entry{e.node.Left, e.depth + 1})
	}
	return maximum
}

func maxDepthIterativePostOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	stack := []entry{{root, 1}}
	var lastVisited *TreeNode
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		node := e.node
		depth := e.depth
		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		if (node.Left == nil && node.Right == nil) || (lastVisited != nil && (lastVisited == node.Left || lastVisited == node.Right)) {
			stack = stack[:len(stack)-1]
			maximum = max(maximum, depth)
			lastVisited = node
		} else {
			stack = append(stack, entry{node.Right, depth + 1})
			stack = append(stack, entry{node.Left, depth + 1})
		}
	}
	return maximum
}

/*
BFS
*/
// pushする前にnilノードを弾く
func maxDepthIterativeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	queue := []entry{{root, 1}}
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		maximum = max(maximum, e.depth)
		if e.node.Left != nil {
			queue = append(queue, entry{e.node.Left, e.depth + 1})
		}
		if e.node.Right != nil {
			queue = append(queue, entry{e.node.Right, e.depth + 1})
		}
	}
	return maximum
}

// popした後にnilノードを弾く
func maxDepthIterativeBFS2(root *TreeNode) int {
	maximum := 0
	queue := []entry{{root, 1}}
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if e.node == nil {
			continue
		}
		maximum = max(maximum, e.depth)
		queue = append(queue, entry{e.node.Left, e.depth + 1})
		queue = append(queue, entry{e.node.Right, e.depth + 1})
	}
	return maximum
}

// depthを保持せずに処理する
func maxDepthIterativeBFS3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maximum := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		maximum++
		count := len(queue)
		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return maximum
}
