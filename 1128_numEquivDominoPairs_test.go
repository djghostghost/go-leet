package main

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {

	pairs := [][]int{{1, 2}, {2, 1}, {1, 2}, {1, 2}, {1, 2}}
	got := numEquivDominoPairs(pairs)
	t.Log(got)
}
