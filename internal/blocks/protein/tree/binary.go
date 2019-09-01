package tree

import "github.com/kamilsk/genome/internal/blocks/protein"

// BinaryTreeNode is a binary tree node.
type BinaryTreeNode struct {
	Parent  *BinaryTreeNode
	Left    *BinaryTreeNode
	Right   *BinaryTreeNode
	Payload protein.T
}
