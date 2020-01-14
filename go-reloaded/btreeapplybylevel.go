package piscine

func BTreeLevelCount(root *TreeNode) int {

	if root == nil {
		return 0
	}
	countleft := BTreeLevelCount(root.Left)
	countright := BTreeLevelCount(root.Right)

	count := countleft
	if count < countright {
		count = countright
	}
	return count + 1
}
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	level := BTreeLevelCount(root)

	for i := 1; i < level+1; i++ {
		bfs(root, i, f)
	}
}

func bfs(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	if level == 1 {
		f(root.Data)
		return
	}

	bfs(root.Left, level-1, f)
	bfs(root.Right, level-1, f)
	return
}
