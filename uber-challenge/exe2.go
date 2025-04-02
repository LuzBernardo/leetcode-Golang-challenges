package uberchallenge

func findTopTwo(wins, draws, scored, conceded []int) []int {
	if len(wins) == 0 {
		return []int{}
	}
	if len(wins) == 1 {
		return []int{0}
	}

	max := func(avoid int) int {
		idx := 0

		for i := 0; i < len(wins); i++ {
			if i == avoid {
				continue
			}
			if (wins[i]*3 + draws[i]) > (wins[idx]*3 + draws[idx]) {
				idx = i
				continue
			}
			if (wins[i]*3 + draws[i]) == (wins[idx]*3 + draws[idx]) {
				if (scored[i] - conceded[i]) > (scored[idx] - conceded[idx]) {
					idx = i
				}
			}
		}

		return idx
	}

	first := max(-1)
	second := max(first)

	return []int{first, second}
}
