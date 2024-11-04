package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Root -> Left -> Right
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

// Left -> Root -> Right
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	InOrderTraversal(root.Right)
}

// Left -> Right -> Root
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Printf("%d ", root.Val)
}

func PreOrderTraversalIterative(root *Node) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", node.Val)
		// 厳密にpre-orderにするなら、まず右の子ノードを追加してから左の子ノードを追加する
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func InOrderTraversalIterative(root *Node) {
	stack := []*Node{}
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", node.Val)
		node = node.Right
	}
}

func PostOrderTraversalIterative(root *Node) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	var lastVisited *Node
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		// 左右の子が処理済み、または子がない場合にノードを取り出す
		if (node.Left == nil && node.Right == nil) ||
			(lastVisited != nil &&
				(lastVisited == node.Left || lastVisited == node.Right)) {
			stack = stack[:len(stack)-1]
			fmt.Printf("%d ", node.Val)
			lastVisited = node
		} else {
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}
}
