package tree

import (
	"errors"
	"strings"
)
const (
	notOp = "NOT"
)

var (
	opMap = map[string]int{
		"AND": 1,
		"OR":  1,
		"NOT": 2,
		"=":   3,
		"<":   3,
		"<=":  3,
		">":   3,
		">=":  3,
	}
	errInvalid = errors.New("invalid expression")
	errUnbalance = errors.New("unbalance parenthesis")
)

func isOp(val string) bool {
	if _, ok := opMap[val]; ok {
		return true
	}
	return false
}

func opPrecendence(val string) int {
	return opMap[val]
}

func isOpenParenthesis(val string) bool {
	return val == "("
}

func isCloseParenthesis(val string) bool {
	return val == ")"
}

/**
The method was implemented based on Shunting-yard algorithm
Wiki: https://en.wikipedia.org/wiki/Shunting-yard_algorithm
Algorithms:
while there are tokens to be read do:
    read a token.
    if the token is a number, then:
        push it to the output queue.
    if the token is a function then:
        push it onto the operator stack 
    if the token is an operator, then:
        while ((there is a function at the top of the operator stack)
               or (there is an operator at the top of the operator stack with greater precedence)
               or (the operator at the top of the operator stack has equal precedence and is left associative))
              and (the operator at the top of the operator stack is not a left parenthesis):
            pop operators from the operator stack onto the output queue.
        push it onto the operator stack.
    if the token is a left paren (i.e. "("), then:
        push it onto the operator stack.
    if the token is a right paren (i.e. ")"), then:
        while the operator at the top of the operator stack is not a left paren:
            pop the operator from the operator stack onto the output queue.
        // if the stack runs out without finding a left paren, then there are mismatched parentheses.
        if there is a left paren at the top of the operator stack, then:
            pop the operator from the operator stack and discard it
after while loop, if operator stack not null, pop everything to output queue
if there are no more tokens to read then:
    while there are still operator tokens on the stack:
        // if the operator token on the top of the stack is a paren, then there are mismatched parentheses.
        pop the operator from the operator stack onto the output queue.
exit.
*/
func infixToTree(tokens []string) (*Node, error) {
	opStk := &stringStack{}
	outStk := &nodeStack{}

	for _, token := range tokens {
		toUpperToken := strings.ToUpper(token)
		if !isOp(toUpperToken) && !isOpenParenthesis(toUpperToken) && !isCloseParenthesis(toUpperToken) {
			outStk.push(&Node{val: token})
			continue
		}

		if isOp(toUpperToken) {
			for opStk.size() != 0 && 
				(isOp(opStk.peek()) && (opPrecendence(opStk.peek()) > opPrecendence(toUpperToken))) &&
				!isOpenParenthesis(opStk.peek()) {
					addNode(outStk, opStk.pop())
			}
			opStk.push(toUpperToken)
			continue
		}

		if isOpenParenthesis(toUpperToken) {
			opStk.push(toUpperToken)
			continue
		}

		if isCloseParenthesis(toUpperToken) {
			for opStk.size() != 0 && !isOpenParenthesis(opStk.peek()) {
				addNode(outStk, opStk.pop())
			}

			if opStk.size() == 0 {
				return nil, errUnbalance
			}

			if isOpenParenthesis(opStk.peek()) {
				opStk.pop()
			}
			continue
		}
	}

	for opStk.size() != 0 && isOp(opStk.peek()) {
		addNode(outStk, opStk.pop())
	}

	if opStk.size() != 0 {
		if isOpenParenthesis(opStk.peek()) {
			return nil, errUnbalance
		}
		return nil, errInvalid
	}

	node := outStk.pop()
	if outStk.size() != 0 {
		return nil, errInvalid
	}

	err := node.Validate()
	if err != nil {
		return nil, err
	}
	return node, nil
}

func addNode(outStk *nodeStack, op string) {
	if op == notOp {
		left := outStk.pop()
		n := &Node{val: op, left: left}
		outStk.push(n)
		return
	}
	right := outStk.pop()
	left := outStk.pop()
	n := &Node{val: op, left: left, right: right}
	outStk.push(n)
}