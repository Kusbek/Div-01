package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var head *NodeI
	var node *NodeI
	var wluha *NodeI
	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}
	if n1.Data < n2.Data {
		head = n1
		wluha = n1
		node = n2
	} else {
		head = n2
		wluha = n2
		node = n1
	}

	for node != nil {
		wluha = SortListInsert(wluha, node.Data)
		node = node.Next
	}
	head = ListSort(head)

	return head
}

func ListSort(node *NodeI) *NodeI {
	node1 := node
	for node1 != nil {
		node2 := node1.Next
		for node2 != nil {
			if node1.Data > node2.Data {
				swap(node1, node2)
			}
			node2 = node2.Next
		}
		node1 = node1.Next
	}
	return node
}

func swap(node1, node2 *NodeI) {
	tmp := node1.Data

	node1.Data = node2.Data
	node2.Data = tmp
}
