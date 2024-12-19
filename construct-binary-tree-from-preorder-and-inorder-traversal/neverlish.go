package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	result := []int{}

	root := preorder[0]
	result = append(result, root)

	preorder = preorder[1:]

	
	// inorder 에서 root 의 위치를 구함
	rootIndex := 0
	for i, v := range inorder {
		if v == root {
			rootIndex = i
			break
		}
	}

	left :
}