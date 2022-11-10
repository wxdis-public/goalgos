package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

// Returns a new tree, initialized with the provided
// values.
func New(values ...int) *Tree {
	tree := &Tree{}
	for i := range values {
		tree.Insert(values[i])
	}
	return tree
}

// Checks if a value exists in the tree, returning
// a pointer to it if it does.
func (t *Tree) Find(value int) *Node {
	if t.Root == nil {
		return nil
	}
	if t.Root.Value == value {
		return t.Root
	}
	return t.Root.find(value)
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{
			Value: value,
		}
		return
	}
	t.Root.insert(value)
}

func (n *Node) find(value int) *Node {
	if n.Value == value {
		return n
	}

	if n.Value < value && n.Right != nil {
		return n.Right.find(value)
	}

	if n.Value > value && n.Left != nil {
		return n.Left.find(value)
	}

	return nil
}

func (n *Node) insert(value int) {
	if n == nil {
		return
	}
	if value == n.Value {
		return
	}
	if n.Value > value {
		if n.Left == nil {
			n.Left = &Node{
				Value: value,
			}
		} else {
			n.Left.insert(value)
		}
		return
	}
	if n.Right == nil {
		n.Right = &Node{
			Value: value,
		}
	} else {
		n.Right.insert(value)
	}

}
