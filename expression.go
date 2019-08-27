package tree

import (
	"fmt"
)

// InfixToExp uses Golang text/scanner and Shunting-yard algorithm to parse expression to tree
func InfixToExp(infixExp string) (*Expression, error) {
	tokens := tokenize(infixExp)

	node, err := infixToTree(tokens)
	if err != nil {
		return nil, err
	}
	return &Expression{
		tree: node,
	}, nil
}

// Expression represents the boolean expression tree and allows basic operation on this
type Expression struct {
	tree *Node
}

// findNode implements DFS on binary tree
func (e *Expression) findNode(key string, predicate func(n *Node) *Node) *Node {
	stk := &nodeStack{}
	stk.push(e.tree)

	for stk.size() != 0 {
		n := stk.pop()
		
		result := predicate(n)
		if result != nil {
			return result
		}

		if n.left != nil {
			stk.push(n.left)
		}
		if n.right != nil {
			stk.push(n.right)
		}
	}
	return nil
}

// FindStatement returns the node which contains the key
func (e *Expression) FindStatement(key string) *Node {
	predicate := func(n *Node) *Node {
		if n.left != nil && n.left.val == key {
			return n
		}

		if isOp(n.val) && 
			(opPrecendence(n.val) == 1 || opPrecendence(n.val) == 2) &&
			(n.right != nil && !isOp(n.right.val) && n.right.val == key) {
				return n.right
		}
		return nil
	}
	return e.findNode(key, predicate)
}

// RemoveStatement finds a key in the expression and remove it from the tree
// return err if key can be found
// the first node will be removed
func (e *Expression) RemoveStatement(key string) error {
	n := e.FindStatement(key)
	if n == nil {
		return fmt.Errorf("%s can't be found", key)
	}
	return nil
}

// ToQuery returns query from the tree
func (e *Expression) ToQuery() string {
	return e.tree.ToQuery()
}
