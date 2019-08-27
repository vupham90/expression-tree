package tree

import (
	"strings"
	"text/scanner"
)

func tokenize(exp string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(exp))

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}