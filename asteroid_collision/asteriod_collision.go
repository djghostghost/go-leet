package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	var result []int

	if len(asteroids) > 0 {
		result = append(result, asteroids[0])
	}

	for i := 1; i < len(asteroids); i++ {
		e := asteroids[i]
		for {
			if len(result) == 0 {
				break
			}
			h := result[len(asteroids)-1]
			// 方向相反
			if e*h < 0 {
				m := absMax(e, h)
				result = result[0 : len(asteroids)-2]
				result = append(result, m)
				if m == e {

				} else {

				}
			}
		}
	}
	for _, e := range asteroids {
		if len(result) == 0 {
			result = append(result, e)
		} else {

		}
	}

	return result
}

func absMax(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a > b {
		return a
	}
	return b
}
