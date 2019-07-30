package main

func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		a, b, c := 0, 1, 1
		for i := 2; i < n; i++ {
			a, b, c = b, c, a+b+c
		}
		return c
	}
}
