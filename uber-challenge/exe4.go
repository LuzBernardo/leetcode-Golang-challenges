package uberchallenge

func countBlackPoints(rows, cols int, blackpoints [][]int) []int {
	results := make([]int, 5)

	if rows == 0 || cols == 0 {
		return results
	}

	mapBlackPoints := make(map[[2]int]bool, 0)

	for _, val := range blackpoints {
		mapBlackPoints[[2]int{val[0], val[1]}] = true
	}

	if rows < 2 || cols < 2 {
		if rows == 1 && cols >= 2 {
			for j := 0; j < cols-2; j++ {
				counter := 0
				for c := j; c < j+1; j++ {
					if mapBlackPoints[[2]int{0, c}] {
						counter++
					}
				}
				results[counter]++
			}
			return results
		}
		if rows >= 2 && cols == 1 {
			for i := 0; i < rows-2; i++ {
				counter := 0
				for r := i; r < i+1; i++ {
					if mapBlackPoints[[2]int{r, 0}] {
						counter++
					}
				}
				results[counter]++
			}
			return results
		}
		if rows == 1 && cols == 1 {
			counter := 0
			if mapBlackPoints[[2]int{0, 0}] {
				counter++
			}
			results[counter]++
			return results
		}
	}

	for i := 0; i < rows-2; i++ {
		for j := 0; j < cols-2; j++ {
			counter := 0
			for r := i; r < r+1; r++ {
				for c := j; c < j+1; j++ {
					if mapBlackPoints[[2]int{r, c}] {
						counter++
					}
				}
			}
			results[counter]++
		}
	}

	return results
}
