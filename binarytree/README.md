# Binary Tree

A Binary Tree is a balanced tree comprised of nodes structured like:

```go
type node struct {
    value int
    left *node
    right *node
}
```

Nodes to the left hand side of the current node are lower than the value of the current node, and nodes to the right of the current node are higher. This diagram demonstrates an example of a binary search tree:

```mermaid
graph TD
    A(5)-->|Left|C
    A-->|Right|B
    C-->|Left|D
    C-->|Right|E
    B-->|Right|F 
    F-->|Right|G
    G-->|Left|I
    G-->|Right|J
    B(8)
    C(3)
    D(1)
    E(4)
    F(11)
    G(15)
    I(14)
    J(18)
```

If the tree exhibits the behaviour above, it is said to be _balanced_ - however if a binary search tree is not balanced, then it shouldn't be considered a binary search tree.