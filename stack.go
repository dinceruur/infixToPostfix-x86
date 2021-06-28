package main

// Stack is used to store operands and operators.
// top points the top of the stack.
// size holds the size of the stack.
type Stack struct {
	top  *Element
	size int
}

// Element is the single member of the Stack: it can be an operand or operator.
// value holds the value of the element, for instance it can be "(" or "+" or "5".
// next points to the subsequent/next element of the stack
type Element struct {
	value string
	next  *Element
}

// Empty checks the stack size and returns true if it is empty
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Top returns the value of the top of the stack
func (s *Stack) Top() string {
	return s.top.value
}

// Push adds the given element to the top of the stack,
// and increases its size by one.
func (s *Stack) Push(value string) {
	s.top = &Element{value, s.top}
	s.size++
}

// Pop returns the top of the stack's value and decreases the size of the stack by one
// And it changes the top of the stack with the next element.
func (s *Stack) Pop() (value string) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value
	}
	return ""
}