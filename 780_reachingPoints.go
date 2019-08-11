package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {

	for true {

		if tx < sx || ty < sy {
			break
		}

		if tx > ty {
			if ty == sy && (tx-sx)%ty == 0 {
				return true
			}
			tx = tx % ty
		} else {

			if tx == sx && (ty-sy)%tx == 0 {
				return true
			}
			ty = ty % tx

		}

	}
	return false
}
