package main

import (
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	s := make(stack, 0)
	s = s.Push("1")
	s = s.Push("2")
	s = s.Push("3")
	for i := 3; i > 0; i-- {
		stack, v, err := s.Pop()
		if err != nil {
			t.Fail()
		}
		s = stack
		converted_value, err := strconv.Atoi(v)
		if err != nil || converted_value != i {
			t.Fail()
		}
	}
	_, _, err := s.Pop()
	if err == nil {
		t.Fatalf("Pop on empty stack should raise an erro")
	}
}
