package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	insNode := &NodeI{Data: data_ref}
	head := l
	node := l
	for node.Next != nil {
		if node.Data >= insNode.Data {
			insNode.Next = node
			head = insNode
			return head
		}
		if node.Data <= insNode.Data && node.Next.Data >= insNode.Data {
			insNode.Next = node.Next
			node.Next = insNode
			return head
		}
		node = node.Next
	}

	node.Next = insNode
	insNode.Next = nil
	return head
}
