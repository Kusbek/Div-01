package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	parent := node.Parent
	if parent == nil {
		min := BTreeMin(node.Right)
		if min == node.Right {
			min.Left = node.Left
			min.Left.Parent = min
			min.Parent = nil
			return min
		}
		nodeRight := BTreeDeleteNode(node.Right, min)
		min.Left = node.Left
		min.Right = nodeRight
		return min
	}
	if node.Left != nil && node.Right != nil {
		min := BTreeMin(node.Right)
		if node.Right == min {
			min.Left = node.Left
			min.Left.Parent = min
			min.Parent = node.Parent
			// fmt.Println("Data: ", min.Data, "           Left: ", min.Left, "             Right: ", min.Right, "              Parent: ", min.Parent)
			BTreeTransplant(root, node, min)
			return root
		}
		nodeRight := BTreeDeleteNode(node.Right, min)
		min.Left = node.Left
		min.Right = nodeRight
		BTreeTransplant(root, node, min)
		return root
	}
	if node.Left == nil && node.Right == nil {
		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return root
	}
	if node.Left == nil || node.Right == nil {
		replace := node.Left
		if replace == nil {
			replace = node.Right
		}
		BTreeTransplant(root, node, replace)
		return root
	}

	return root
}
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}
