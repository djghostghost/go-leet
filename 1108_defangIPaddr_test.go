package main

import "testing"

func Test_defangIPaddr(t *testing.T) {

	t.Log(defangIPaddr("255.100.50.0"))
}
