package pointers

type LinkedList interface {
	Push(data interface{}) error
	Get(index int) (*Node, error)
	Insert(data interface{}, position int) error
	Remove(index int) (*Node, error)
	Search(data interface{}) (*Node, error)
	Size() (int, error)
}

type linkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{nil}
}

func (ll *linkedList) Push(data interface{}) error {
	return nil
}
func (ll *linkedList) Get(index int) (*Node, error) {
	return nil, nil
}
func (ll *linkedList) Insert(data interface{}, position int) error {
	return nil
}
func (ll *linkedList) Remove(index int) (*Node, error) {
	return nil, nil
}
func (ll *linkedList) Search(data interface{}) (*Node, error) {
	return nil, nil
}
func (ll *linkedList) Size() (int, error) {
	return 0, nil
}
