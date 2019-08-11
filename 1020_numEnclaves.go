package main

func numEnclaves(A [][]int) int {

	m := len(A)
	n := len(A[0])

	union := NewUnionFind(m * n)

	mark := make(map[int]bool)

	DIRS := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if (i == 0 || j == 0 || i == m-1 || j == n-1) && A[i][j] == 1 {
				mark[i*n+j] = true
			}

			if A[i][j] == 1 {
				for _, dir := range DIRS {

					x := i + dir[0]
					y := j + dir[1]

					if x >= 0 && x < m && y >= 0 && y < n && A[x][y] == 1 {
						union.Union(i*n+j, x*n+y)
					}
				}
			}

		}
	}

	counter := make(map[int]int)

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			if A[i][j] == 1 {
				counter[union.Find(i*n+j)]++
			}
		}
	}

	for key := range mark {
		counter[union.Find(key)] = 0
	}
	result := 0

	for _, value := range counter {
		result += value
	}

	return result
}

type UnionFind struct {
	unionSet []int
}

func NewUnionFind(n int) *UnionFind {

	set := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = i
	}

	return &UnionFind{unionSet: set}
}

func (u *UnionFind) Union(a int, b int) {

	aParent := u.Find(a)
	bParent := u.Find(b)

	if aParent == bParent {
		return
	}
	u.unionSet[aParent] = bParent
}

func (u *UnionFind) Find(a int) int {

	if u.unionSet[a] != a {
		u.unionSet[a] = u.Find(u.unionSet[a])
	}
	return u.unionSet[a]
}
