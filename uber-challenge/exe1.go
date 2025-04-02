package uberchallenge

func count(numbs []int, pivot int) string {
	if len(numbs) == 0 {
		return "tie"
	}

	counterGreater := 0
	counterLess := 0

	for _, numb := range numbs {
		if pivot > numb {
			counterGreater++
		}
		if pivot < numb {
			counterLess++
		}
	}

	if counterGreater > counterLess {
		return "greater"
	}
	if counterGreater < counterLess {
		return "smaller"
	}

	return "tie"
}
