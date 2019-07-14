package bt

func generateParenthesis(n int) []string {

	result := make([]string, 0)

	helper(n, 1, 0, "（", &result)
	return result
}

func helper(n int, lcount int, rcount int, combination string, result *[]string) {

	if lcount > n || rcount > n || lcount < rcount {
		return
	}

	if lcount == n && lcount == rcount {
		*result = append(*result, combination)
		return
	}

	helper(n, lcount+1, rcount, combination+"(", result)
	helper(n, lcount, rcount+1, combination+")", result)
}
