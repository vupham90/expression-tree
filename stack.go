package tree

type stringStack struct {
	data []string
}

func (s *stringStack) pop() string {
	size := s.size()
	if size == 0 {
		return ""
	}
	val := s.data[size-1]
	s.data = s.data[0:size-1]
	return val
}

func (s *stringStack) size() int {
	return len(s.data)
}

func (s *stringStack) peek() string {
	if s.size() == 0 {
		return ""
	}
	return s.data[s.size() - 1]
}

func (s *stringStack) push(val string) {
	s.data = append(s.data, val)
}


type nodeStack struct {
	data []*Node
}

func (s *nodeStack) pop() *Node {
	size := s.size()
	if size == 0 {
		return nil
	}
	val := s.data[size-1]
	s.data = s.data[0:size-1]
	return val
}

func (s *nodeStack) size() int {
	return len(s.data)
}

func (s *nodeStack) peek() *Node {
	if s.size() == 0 {
		return nil
	}
	return s.data[s.size() - 1]
}

func (s *nodeStack) push(val *Node) {
	s.data = append(s.data, val)
}