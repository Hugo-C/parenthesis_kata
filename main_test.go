package main

import (
	"testing"
)

func TestSingleParanthesis(t *testing.T) {
	is_valid := IsParenthesisValid("(wow)")
	if !is_valid {
		t.Fail()
	}
}

func TestParanthesisNotClosedInvalid(t *testing.T) {
	is_valid := IsParenthesisValid("(wow")
	if is_valid {
		t.Fail()
	}
}

func TestParanthesisNotMatchingInvalid(t *testing.T) {
	is_valid := IsParenthesisValid("(wow]")
	if is_valid {
		t.Fail()
	}
}

func TestMixedParanthesis(t *testing.T) {
	is_valid := IsParenthesisValid("[{(wow)}]")
	if !is_valid {
		t.Fail()
	}
}

func TestMixedParanthesisInvalid(t *testing.T) {
	is_valid := IsParenthesisValid("[(wow)}]")
	if is_valid {
		t.Fail()
	}
}

func TestComplexParanthesis(t *testing.T) {
	is_valid := IsParenthesisValid("[()({wow})]")
	if !is_valid {
		t.Fail()
	}
}

func TestOverlappingParanthesisInvalid(t *testing.T) {
	is_valid := IsParenthesisValid("[([)]]")
	if is_valid {
		t.Fail()
	}
}
