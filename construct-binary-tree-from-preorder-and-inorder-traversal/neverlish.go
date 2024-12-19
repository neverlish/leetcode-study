package main

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}

	result := buildTree(preorder, inorder)

	if result.Val != 1 {
		t.Errorf("Expected 1, got %d", result.Val)
	}

	if result.Left.Val != 2 {
		t.Errorf("Expected 2, got %d", result.Left.Val)
	}

	if result.Right != nil {
		t.Errorf("Expected nil, got %v", result.Right)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
//         if not inorder:
//             return None

//         val = preorder.pop(0)  # O(n)
//         mid = inorder.index(val)  # O(n)
//         left = self.buildTree(preorder, inorder[:mid])
//         right = self.buildTree(preorder, inorder[mid + 1 :])
//         return TreeNode(val, left, right)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	val := preorder[0]
	preorder = preorder[1:]

	mid := 0
	for i, v := range inorder {
		if v == val {
			mid = i
			break
		}
	}

	left := buildTree(preorder, inorder[:mid])
	right := buildTree(preorder, inorder[mid+1:])

	return &TreeNode{val, left, right}
}
