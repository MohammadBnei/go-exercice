package pointers

type Node struct {
	Data     interface{}
	NextNode *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}
