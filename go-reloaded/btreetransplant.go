package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	parent := node.Parent
	if parent == nil {
		return rplc
	}
	if parent.Left == node {
		parent.Left = rplc
		rplc.Parent = parent
		return root
	}

	if parent.Right == node {
		parent.Right = rplc
		rplc.Parent = parent
		return root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	}
	return BTreeSearchItem(root.Left, elem)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data, Parent: root}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f("data: ", root.Data, ", parent: ", &root.Parent, "address: ", &root)
	BTreeApplyInorder(root.Right, f)
}
