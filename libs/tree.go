package libs

import (
	"bytes"
	"fmt"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func binaryTreeToString(root *BinaryTreeNode) string {
	if root == nil {
		return "()"
	}

	left := binaryTreeToString(root.Left)
	right := binaryTreeToString(root.Right)

	var result bytes.Buffer

	result.WriteString(fmt.Sprintf("(%d", root.Val))

	if len(left) > 2 || len(right) > 2 {
		result.WriteString(fmt.Sprintf(" %s %s", left, right))
	}

	result.WriteString(")")

	return result.String()
}

func (this *BinaryTreeNode) String() string {
	return binaryTreeToString(this)
}

func DfsInOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	var nums []int = []int{}

	nums = append(nums, DfsInOrder(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, DfsInOrder(root.Right)...)

	return nums
}

func DfsPreOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	var nums []int = []int{}

	nums = append(nums, root.Val)
	nums = append(nums, DfsPreOrder(root.Left)...)
	nums = append(nums, DfsPreOrder(root.Right)...)

	return nums
}

func DfsPostOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	var nums []int = []int{}

	nums = append(nums, DfsPostOrder(root.Left)...)
	nums = append(nums, DfsPostOrder(root.Right)...)
	nums = append(nums, root.Val)

	return nums
}

// 使用一个队列，然后从root节点开始，每处理一个节点，就将该节点的左、右节点放入队列尾部。
func Bfs(root *BinaryTreeNode) []int {
	result := []int{}

	BfsTraversal(root, func(node *BinaryTreeNode) bool {
		result = append(result, node.Val)
		return false
	})

	return result
}

func BfsTraversal(root *BinaryTreeNode, f func(*BinaryTreeNode) bool) {
	if root == nil {
		return
	}

	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		if f(node) {
			return
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		queue = queue[1:]
	}
}

func ReverseBinaryTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return root
	}

	left := ReverseBinaryTree(root.Right)
	right := ReverseBinaryTree(root.Left)

	root.Left = left
	root.Right = right

	return root
}

func BfsInsert(root *BinaryTreeNode, key int) *BinaryTreeNode {
	keyNode := &BinaryTreeNode{
		Val: key,
	}
	if root == nil {
		return keyNode
	}

	BfsTraversal(root, func(node *BinaryTreeNode) bool {
		if node.Left == nil {
			node.Left = keyNode
			return true
		}

		if node.Right == nil {
			node.Right = keyNode
			return true
		}

		return false
	})

	return root
}

func DeleteFromBinaryTree(root *BinaryTreeNode, key int) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	var parentNode *BinaryTreeNode
	rightMostNode := root
	newRoot := root

	if root.Right == nil {
		newRoot = root.Left
	}

	for rightMostNode.Right != nil {
		parentNode = rightMostNode
		rightMostNode = rightMostNode.Right
	}

	BfsTraversal(root, func(node *BinaryTreeNode) bool {
		if node.Val == key {
			node.Val = rightMostNode.Val
			if parentNode != nil {
				parentNode.Right = rightMostNode.Left
			}
			return true
		}
		return false
	})

	return newRoot
}

// BST
type BSTNode struct {
	Val     int
	counter uint
	Left    *BSTNode
	Right   *BSTNode
}

func bstTreeToString(root *BSTNode) string {
	if root == nil {
		return "()"
	}

	left := bstTreeToString(root.Left)
	right := bstTreeToString(root.Right)

	var result bytes.Buffer

	result.WriteString(fmt.Sprintf("(%d", root.Val))

	if root.counter > 1 {
		result.WriteString(fmt.Sprintf("/%d", root.counter))
	}

	if len(left) > 2 || len(right) > 2 {
		result.WriteString(fmt.Sprintf(" %s %s", left, right))
	}

	result.WriteString(")")

	return result.String()
}

func (this *BSTNode) String() string {
	return bstTreeToString(this)
}

func BSTInsert(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return &BSTNode{
			counter: 1,
			Val:     val,
		}
	}

	if root.Val == val {
		root.counter++
	} else if root.Val > val {
		root.Left = BSTInsert(root.Left, val)
	} else {
		root.Right = BSTInsert(root.Right, val)
	}
	return root
}

func BSTDelete(root *BSTNode, val int) *BSTNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		root.counter--

		if root.counter == 0 {
			if root.Right == nil {
				return root.Left
			}

			// 寻找后继节点，后继节点是右子树中最小的那个节点
			parent := root
			node := root.Right
			for node.Left != nil {
				parent = node
				node = node.Left
			}

			// 如果后继节点就是当前根节点的右节点，则将其提升上去即可
			if parent == root {
				node.Left = root.Left
				return node
			}

			// 如果后继节点在其它地方，则需要处理该后继节点的父节点与其右子树的关系（左子树为nil）
			parent.Left = node.Right
			node.Right = root.Right
			node.Left = root.Left

			return node
		} else {
			return root
		}
	} else if root.Val > val {
		root.Left = BSTDelete(root.Left, val)
	} else {
		root.Right = BSTDelete(root.Right, val)
	}

	return root
}

////////////////////////////////////////////////
//     y                               x	   //
//    / \     Right Rotation          /  \	   //
//   x   T3   - - - - - - - >        T1   y   //
//  / \       < - - - - - - -            / \  //
// T1  T2     Left Rotation            T2  T3 //
////////////////////////////////////////////////
func LeftRotation(root *BSTNode) *BSTNode {
	if root == nil || root.Right == nil {
		return root
	}

	right := root.Right

	root.Right = right.Left
	right.Left = root

	return right
}

func RightRatation(root *BSTNode) *BSTNode {
	if root == nil || root.Left == nil {
		return root
	}

	left := root.Left

	root.Left = left.Right
	left.Right = root

	return left
}

func BSTHeight(root *BSTNode) int {
	if root == nil {
		return -1
	}

	leftHeight := BSTHeight(root.Left)
	rightHeight := BSTHeight(root.Right)

	if leftHeight > rightHeight {
		return 1 + leftHeight
	}

	return 1 + rightHeight
}
