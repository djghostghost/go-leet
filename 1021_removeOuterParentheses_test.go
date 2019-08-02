package main

import "testing"

func Test_removeOuterParentheses(t *testing.T) {

	t.Log(removeOuterParentheses("(()())(())(()(()))"))
}
