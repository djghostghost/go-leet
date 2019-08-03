package main

import (
	"fmt"
	"testing"
)

func Test_recoverFromPreorder(t *testing.T) {

	root := recoverFromPreorder("1-401--349---90--88")

	fmt.Println(root)
}
