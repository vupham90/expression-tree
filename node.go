package tree

import (
	"errors"
	"strings"
)

type Node struct {
	val string
	left *Node
	right *Node
}

func (n *Node) ToQuery() string {
	query := ""

	if n.val != notOp {
		if n.left != nil {
			query += "(" + n.left.ToQuery() + " "
		}
		query += n.val + " "
		if n.right != nil {
			query += n.right.ToQuery() + ") "
		}
	} else {
		query += n.val + " "
		if n.left != nil {
			query += "(" + n.left.ToQuery() + ") "
		}
	}
	
	return strings.Trim(query, " ")
}

func (n *Node) Validate() error {
	if isOp(n.val) {
		if n.val == notOp {
			if n.left == nil || n.right != nil {
				return errors.New("invalid expression for " + n.val)
			}
			leftValidate := n.left.Validate()
			if leftValidate != nil {
				return leftValidate
			}
			return nil
		}

		if n.left == nil || n.right == nil {
			return errors.New("invalid expression for " + n.val)
		}

		leftValidate := n.left.Validate()
		if leftValidate != nil {
			return leftValidate
		}

		rightValidate := n.right.Validate()
		if rightValidate != nil {
			return rightValidate
		}
	}
	return nil
}