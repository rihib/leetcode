//lint:file-ignore U1000 Ignore all unused code
package main

import "container/list"

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

func sortedArrayToBSTClosed(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTClosedHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTClosedHelper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBSTClosedHelper(nums, left, mid-1),
		Right: sortedArrayToBSTClosedHelper(nums, mid+1, right),
	}
}

func sortedArrayToBSTHalfOpen(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTHalfOpenHelper(nums, 0, len(nums))
}

func sortedArrayToBSTHalfOpenHelper(nums []int, left, right int) *TreeNode {
	if left == right {
		return nil
	}
	mid := left + (right-left)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBSTHalfOpenHelper(nums, left, mid),
		Right: sortedArrayToBSTHalfOpenHelper(nums, mid+1, right),
	}
}

type bstElement struct {
	root  *TreeNode
	left  []int
	right []int
}

func sortedArrayToBSTDFS(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	stack := []bstElement{{root, nums[:midIndex], nums[midIndex+1:]}}
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(e.left) > 0 {
			leftMid := len(e.left) / 2
			e.root.Left = &TreeNode{Val: e.left[leftMid]}
			stack = append(stack, bstElement{
				e.root.Left, e.left[:leftMid], e.left[leftMid+1:],
			})
		}
		if len(e.right) > 0 {
			rightMid := len(e.right) / 2
			e.root.Right = &TreeNode{Val: e.right[rightMid]}
			stack = append(stack, bstElement{
				e.root.Right, e.right[:rightMid], e.right[rightMid+1:],
			})
		}
	}
	return root
}

func sortedArrayToBSTBFS(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}
	queue := list.New()
	queue.PushBack(bstElement{root, nums[:midIndex], nums[midIndex+1:]})
	for queue.Len() > 0 {
		e := queue.Remove(queue.Front()).(bstElement)
		if len(e.left) > 0 {
			leftMid := len(e.left) / 2
			e.root.Left = &TreeNode{Val: e.left[leftMid]}
			queue.PushBack(bstElement{
				e.root.Left, e.left[:leftMid], e.left[leftMid+1:],
			})
		}
		if len(e.right) > 0 {
			rightMid := len(e.right) / 2
			e.root.Right = &TreeNode{Val: e.right[rightMid]}
			queue.PushBack(bstElement{
				e.root.Right, e.right[:rightMid], e.right[rightMid+1:],
			})
		}
	}
	return root
}

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
