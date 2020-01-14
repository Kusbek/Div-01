package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := NodeL{Data: data}
	if l.Head == nil {
		l.Head = &newNode
	} else {
		for node := l.Head; node != nil; node = node.Next {
			if node.Next == nil {
				node.Next = &newNode
				return
			}
		}
	}
}
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
		ListRemoveIf(l, data_ref)
		return
	}

	node := l.Head

	for node.Next != nil {
		if node.Next.Data == data_ref {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
}
