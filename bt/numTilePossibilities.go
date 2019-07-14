package bt

func numTilePossibilities(tiles string) int {

	arry := make([]int, 26)

	for _, ch := range tiles {

		arry[ch-'A']++

	}

	return backtracking(&arry)
}

func backtracking(array *[]int) int {

	total := 0

	for i := range *array {

		if (*array)[i] == 0 {
			continue
		} else {
			(*array)[i] -= 1
			total++
			total += backtracking(array)
			(*array)[i] += 1
		}
	}

	return total
}
