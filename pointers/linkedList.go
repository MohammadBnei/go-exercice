package pointers

type LinkedList interface {
	Push(data interface{}) error
	Get(index int) (*Node, error)
	Insert(data interface{}, position int) error
	Remove(index int) (*Node, error)
	Search(data interface{}) (*Node, error)
	Size() (int, error)
}
