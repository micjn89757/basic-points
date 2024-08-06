package tree 

import "testing"


func TestLC133(t *testing.T) {
	root := &BinaryTreeNode{
		Data: 1,
		LeftNode: &BinaryTreeNode{
			Data: 2,
			LeftNode: nil,
			RightNode: nil,
		},
		RightNode: &BinaryTreeNode{
			Data: 3,
			LeftNode: nil,
			RightNode: nil,
		},
	}

	expect := 3

	t.Logf("%#v", pathSum(root, expect))
}