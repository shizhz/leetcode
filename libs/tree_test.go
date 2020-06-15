package libs

import (
	"fmt"
	"reflect"
	"testing"
)

func makeTestBinaryTree() *BinaryTreeNode {
	return &BinaryTreeNode{
		Val: 1,
		Left: &BinaryTreeNode{
			Val: 2,
			Left: &BinaryTreeNode{
				Val: 4,
			},
			Right: &BinaryTreeNode{
				Val: 5,
			},
		},
		Right: &BinaryTreeNode{
			Val: 3,
		},
	}
}

func TestDfsInOrder(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want []int
	}{
		{
			name: "test 01",
			root: nil,
			want: nil,
		},
		{
			name: "test 02",
			root: makeTestBinaryTree(),
			want: []int{4, 2, 5, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DfsInOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DfsInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDfsPreOrder(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want []int
	}{
		{
			name: "test 01",
			root: nil,
			want: nil,
		},
		{
			name: "test 02",
			root: makeTestBinaryTree(),
			want: []int{1, 2, 4, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DfsPreOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DfsInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDfsPostOrder(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want []int
	}{
		{
			name: "test 01",
			root: nil,
			want: nil,
		},
		{
			name: "test 02",
			root: makeTestBinaryTree(),
			want: []int{4, 5, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DfsPostOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DfsInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBfs(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want []int
	}{
		{
			name: "test 01",
			root: makeTestBinaryTree(),
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test 02",
			root: nil,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bfs(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseBinaryTree(t *testing.T) {
	root := makeTestBinaryTree()
	fmt.Println("OriginalTree: ", root)
	fmt.Println("ReversedTree: ", ReverseBinaryTree(root))
	fmt.Println("BfsInsert: ", BfsInsert(root, 10))
	fmt.Println("DeleteFromBinaryTree: ", DeleteFromBinaryTree(makeTestBinaryTree(), 1))
	fmt.Println("DeleteFromBinaryTree: ", DeleteFromBinaryTree(makeTestBinaryTree(), 2))
	fmt.Println("DeleteFromBinaryTree: ", DeleteFromBinaryTree(makeTestBinaryTree(), 3))
	fmt.Println("DeleteFromBinaryTree: ", DeleteFromBinaryTree(DeleteFromBinaryTree(makeTestBinaryTree(), 3), 1))
}

func TestBST(t *testing.T) {
	root := BSTInsert(nil, 1)
	root = BSTInsert(root, -2)
	root = BSTInsert(root, 2)
	root = BSTInsert(root, 2)
	root = BSTInsert(root, 8)
	root = BSTInsert(root, 7)

	fmt.Println(root)
	fmt.Printf("Height: %d\n", root.height())
	fmt.Println("Test BST Deletion")
	root = BSTDelete(root, 2)
	fmt.Println(root)
	root = BSTDelete(root, 7)
	fmt.Println(root)
	root = BSTDelete(root, 1)
	fmt.Println(root)
	root = BSTDelete(root, 8)
	fmt.Println(root)
	root = BSTDelete(root, -2)
	fmt.Println(root)
	root = BSTDelete(root, 2)
	fmt.Println(root)
	root = BSTDelete(root, 2)
	fmt.Println(root)
}

func testAvlInsertByOrder(root *BSTNode, verbose bool, vals ...int) *BSTNode {
	if verbose {
		fmt.Println("==============================")
	}
	for _, val := range vals {
		root = AvlInsert(root, val)
		if verbose {
			fmt.Printf("Inserted %d: %s\n", val, root.String())
		}
	}

	return root
}

func testAvlDeleteByOrder(root *BSTNode, verbose bool, vals ...int) *BSTNode {
	if verbose {
		fmt.Println("==============================")
	}
	for _, val := range vals {
		root = AvlDelete(root, val)
		if verbose {
			fmt.Printf("Deleted %d: %s\n", val, root.String())
		}
	}

	return root
}

func TestAvlTree(t *testing.T) {
	testAvlInsertByOrder(nil, true, 1, -2, 2, 7, 9, 10)
	testAvlInsertByOrder(nil, true, 1, 2, 3, 4, 5, 6)
	testAvlInsertByOrder(nil, true, 1, 10, 8, 7, 5, 6)
	root := testAvlInsertByOrder(nil, true, 1, -10, 10, -15, -5, -3)

	fmt.Println("Test AVLDelete....")

	testAvlDeleteByOrder(root, true, -15, -10)
}	
